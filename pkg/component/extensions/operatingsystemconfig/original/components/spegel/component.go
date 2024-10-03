package spegel

import (
	_ "embed"
	"log"
	"net"
	"strings"

	"k8s.io/utils/ptr"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/component/extensions/operatingsystemconfig/original/components"
	"github.com/gardener/gardener/pkg/component/extensions/operatingsystemconfig/original/components/containerd"
	"github.com/gardener/gardener/pkg/component/extensions/operatingsystemconfig/original/components/kubelet"
	"github.com/gardener/gardener/pkg/utils"
)

// UnitName is the name of the spegel service.
const UnitName = v1beta1constants.OperatingSystemConfigUnitNameSpegelService

//go:embed scripts/metrics.sh
var spegelMetricsContent string

type component struct{}

// New returns a new kubelet component.
func New() *component {
	return &component{}
}

func (component) Name() string {
	return "spegel"
}

func (component) Config(ctx components.Context) ([]extensionsv1alpha1.Unit, []extensionsv1alpha1.File, error) {
	spegelFiles := []extensionsv1alpha1.File{
		{
			Path:        v1beta1constants.OperatingSystemConfigFilePathBinaries + "/spegel",
			Permissions: ptr.To[uint32](0755),
			Content: extensionsv1alpha1.FileContent{
				ImageRef: &extensionsv1alpha1.FileContentImageRef{
					Image:           "ghcr.io/spegel-org/spegel:v0.0.27",
					FilePathInImage: "/app/spegel",
				},
			},
		},
		{
			Path:        v1beta1constants.OperatingSystemConfigFilePathBinaries + "/spegel_metrics.sh",
			Permissions: ptr.To[uint32](0755),
			Content: extensionsv1alpha1.FileContent{
				Inline: &extensionsv1alpha1.FileContentInline{
					Encoding: "b64",
					Data:     utils.EncodeBase64([]byte(spegelMetricsContent)),
				},
			},
		},
	}

	cliFlags := CLIFlags()
	// spegel depends on containerd.sock
	spegelUnit := extensionsv1alpha1.Unit{
		Name:    "spegel.service",
		Command: ptr.To(extensionsv1alpha1.CommandStart),
		Enable:  ptr.To(true),
		Content: ptr.To(`[Unit]
Description=spegel daemon
Documentation=https://github.com/spegel-org/spegel
After=` + containerd.UnitName + `
Requires=` + containerd.UnitName + `
Before=` + kubelet.UnitName + `
[Install]
WantedBy=multi-user.target
[Service]
Restart=always
RestartSec=5
Environment="NODE_IP=` + GetOutboundIP().String() + `"
ExecStart=` + v1beta1constants.OperatingSystemConfigFilePathBinaries + `/spegel \
    ` + utils.Indent(strings.Join(cliFlags, " \\\n"), 4) + "\n"),
		FilePaths: []string{v1beta1constants.OperatingSystemConfigFilePathBinaries + "/spegel"},
	}

	spegelMetricsUnit := extensionsv1alpha1.Unit{
		Name:    "spegel-metrics.service",
		Command: ptr.To(extensionsv1alpha1.CommandStart),
		Enable:  ptr.To(true),
		Content: ptr.To(`[Unit]
Description=spegel metrics daemon
Documentation=https://github.com/spegel-org/spegel
After=` + UnitName + `
BindsTo=` + UnitName + `
[Install]
WantedBy=multi-user.target
[Service]
Restart=always
RestartSec=5
ExecStart=` + v1beta1constants.OperatingSystemConfigFilePathBinaries + `/spegel_metrics.sh`),
		FilePaths: []string{v1beta1constants.OperatingSystemConfigFilePathBinaries + "/spegel_metrics.sh"},
	}

	return []extensionsv1alpha1.Unit{spegelUnit, spegelMetricsUnit}, spegelFiles, nil
}

// Get preferred outbound ip of this machine - TODO
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
