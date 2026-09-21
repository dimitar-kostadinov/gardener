// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package botanist

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resourcesv1alpha1 "github.com/gardener/gardener/pkg/apis/resources/v1alpha1"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/utils/flow"
	kubernetesutils "github.com/gardener/gardener/pkg/utils/kubernetes"
)

// RequiredCleanupsTaskGroup builds the flow.TaskGroup that cleans up the stale resources restored from the ETCD
// snapshot during `gardenadm restore`: it finalizes and deletes the restored ManagedResources, deletes the stale
// gardener-node-agent OperatingSystemConfig Secret, and force-deletes the prior control plane Node (priorNodeName)
// together with the Pods running on it.
//
// clientSet is a pointer to the control plane client set, which is not yet initialized while the graph is being built;
// the tasks must only dereference it from within their Fn (i.e. at flow run time), after the connection to the
// Kubernetes control plane has been established.
func (b *GardenadmBotanist) RequiredCleanupsTaskGroup(clientSet *kubernetes.Interface, priorNodeName string) flow.TaskGroup {
	group := flow.NewTaskGroup("requiredCleanups")

	// FinalizeManagedResources must run before DeleteStaleOperatingSystemConfigSecret: the OperatingSystemConfig
	// Secret is managed by the `shoot-gardener-node-agent` ManagedResource, so that ManagedResource's finalizers
	// must be gone before the Secret can be deleted and MigrateSecrets can reinstall the bootstrap-content Secret
	// under the same name.
	finalizeManagedResources := group.Add(flow.Task{
		Name: "Finalizing and deleting ManagedResources restored from the ETCD snapshot",
		Fn: func(ctx context.Context) error {
			return b.FinalizeManagedResources(ctx, (*clientSet).Client())
		},
	})
	group.Add(flow.Task{
		Name: "Deleting stale gardener-node-agent OperatingSystemConfig Secret restored from the ETCD snapshot",
		Fn: func(ctx context.Context) error {
			return b.DeleteStaleOperatingSystemConfigSecret(ctx, (*clientSet).Client())
		},
		Dependencies: flow.NewTaskIDs(finalizeManagedResources),
	})
	// Deleting the prior control plane Node and its Pods is independent of the ManagedResource/Secret cleanup.
	group.Add(flow.Task{
		Name: "Deleting the prior control plane Node and the Pods running on it",
		Fn: func(ctx context.Context) error {
			return b.DeletePriorNodeAndPodsRunningOnIt(ctx, (*clientSet).Client(), priorNodeName)
		},
	})

	return group
}

// DeleteStaleOperatingSystemConfigSecret deletes the gardener-node-agent OperatingSystemConfig Secret from the given
// (real) cluster client during the restore phase. The Secret restored from the ETCD snapshot carries the *managed*
// etcd static-pod manifests under the content-independent gardener-node-agent Secret name (the name is computed by
// operatingsystemconfig.Key() from cluster parameters, not from the static-pod file list). Deleting it lets the
// subsequent MigrateSecrets task install the *bootstrap*-content OperatingSystemConfig Secret (built during bootstrap
// on the fake seed client) under that same name, so gardener-node-agent applies the bootstrap etcd first - mirroring
// the healthy lineage of `gardenadm init` and the first restore. The later etcd-druid transition then rewrites the
// Secret with managed content and gardener-node-agent removes the bootstrap manifests as usual.
func (b *GardenadmBotanist) DeleteStaleOperatingSystemConfigSecret(ctx context.Context, realClient client.Client) error {
	if b.operatingSystemConfigSecret == nil {
		return fmt.Errorf("operating system config secret is nil, make sure to call createOperatingSystemConfigSecretForNodeAgent() first")
	}

	secret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{
		Name:      b.operatingSystemConfigSecret.Name,
		Namespace: b.operatingSystemConfigSecret.Namespace,
	}}
	return client.IgnoreNotFound(realClient.Delete(ctx, secret))
}

// FinalizeManagedResources removes the finalizers from and deletes all ManagedResources restored from the ETCD
// snapshot, and waits until they are gone. During bootstrap gardener-resource-manager is not yet running on the
// restored control plane, so the finalizers have to be removed explicitly for the ManagedResources to be deleted.
// This must run before DeleteStaleOperatingSystemConfigSecret: the gardener-node-agent OperatingSystemConfig Secret is
// managed by the `shoot-gardener-node-agent` ManagedResource, so its finalizers must be gone before the Secret can be
// deleted.
func (b *GardenadmBotanist) FinalizeManagedResources(ctx context.Context, realClient client.Client) error {
	managedResourceList := &resourcesv1alpha1.ManagedResourceList{}
	if err := realClient.List(ctx, managedResourceList); err != nil {
		return fmt.Errorf("failed listing ManagedResources: %w", err)
	}

	for _, managedResource := range managedResourceList.Items {
		obj := managedResource.DeepCopy()
		obj.SetFinalizers(nil)

		b.Logger.Info("Removing ManagedResource finalizers", "managedResource", client.ObjectKeyFromObject(obj))
		if err := realClient.Update(ctx, obj); client.IgnoreNotFound(err) != nil {
			return fmt.Errorf("failed updating ManagedResource %s: %w", client.ObjectKeyFromObject(obj), err)
		}

		b.Logger.Info("Deleting ManagedResource", "managedResource", client.ObjectKeyFromObject(obj))
		if err := realClient.Delete(ctx, obj); client.IgnoreNotFound(err) != nil {
			return fmt.Errorf("failed deleting ManagedResource %s: %w", client.ObjectKeyFromObject(obj), err)
		}
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	b.Logger.Info("Waiting for ManagedResources to be cleaned up")
	if err := kubernetesutils.WaitUntilResourcesDeleted(ctxWithTimeout, realClient, managedResourceList, 10*time.Second); err != nil {
		return fmt.Errorf("failed waiting until ManagedResources are cleaned up: %w", err)
	}

	return nil
}

// DeletePriorNodeAndPodsRunningOnIt deletes the prior control plane Node (the one being replaced during restore) and
// force-deletes all Pods that were scheduled onto it, so the restored control plane does not keep referencing the lost
// node.
func (b *GardenadmBotanist) DeletePriorNodeAndPodsRunningOnIt(ctx context.Context, realClient client.Client, priorNodeName string) error {
	if priorNodeName == "" {
		// Guard against an empty node name: it would delete a Node named "" and, since unscheduled Pods have an empty
		// .spec.NodeName, force-delete every Pending Pod in the cluster. Callers must pass the prior node name.
		return fmt.Errorf("priorNodeName must not be empty")
	}

	node := &corev1.Node{ObjectMeta: metav1.ObjectMeta{Name: priorNodeName}}

	b.Logger.Info("Deleting Node", "node", client.ObjectKeyFromObject(node))
	if err := realClient.Delete(ctx, node); client.IgnoreNotFound(err) != nil {
		return fmt.Errorf("failed deleting Node %s: %w", client.ObjectKeyFromObject(node), err)
	}

	podList := &corev1.PodList{}
	if err := realClient.List(ctx, podList); err != nil {
		return fmt.Errorf("failed listing Pods: %w", err)
	}

	for _, pod := range podList.Items {
		if pod.Spec.NodeName != priorNodeName {
			continue
		}

		b.Logger.Info("Force deleting Pod", "pod", client.ObjectKeyFromObject(&pod), "nodeName", pod.Spec.NodeName)
		options := &client.DeleteOptions{GracePeriodSeconds: ptr.To[int64](0), PropagationPolicy: ptr.To(metav1.DeletePropagationBackground)}
		if err := realClient.Delete(ctx, pod.DeepCopy(), options); client.IgnoreNotFound(err) != nil {
			return fmt.Errorf("failed force deleting Pod %s: %w", client.ObjectKeyFromObject(&pod), err)
		}
	}

	return nil
}
