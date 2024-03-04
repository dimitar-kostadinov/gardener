// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/component/kubernetes/apiserver (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -package mock -destination=mocks.go github.com/gardener/gardener/pkg/component/kubernetes/apiserver Interface
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	apiserver "github.com/gardener/gardener/pkg/component/apiserver"
	apiserver0 "github.com/gardener/gardener/pkg/component/kubernetes/apiserver"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AlertingRules mocks base method.
func (m *MockInterface) AlertingRules() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertingRules")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertingRules indicates an expected call of AlertingRules.
func (mr *MockInterfaceMockRecorder) AlertingRules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertingRules", reflect.TypeOf((*MockInterface)(nil).AlertingRules))
}

// Deploy mocks base method.
func (m *MockInterface) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockInterfaceMockRecorder) Deploy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockInterface)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockInterface) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockInterfaceMockRecorder) Destroy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockInterface)(nil).Destroy), arg0)
}

// GetAutoscalingReplicas mocks base method.
func (m *MockInterface) GetAutoscalingReplicas() *int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutoscalingReplicas")
	ret0, _ := ret[0].(*int32)
	return ret0
}

// GetAutoscalingReplicas indicates an expected call of GetAutoscalingReplicas.
func (mr *MockInterfaceMockRecorder) GetAutoscalingReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoscalingReplicas", reflect.TypeOf((*MockInterface)(nil).GetAutoscalingReplicas))
}

// GetValues mocks base method.
func (m *MockInterface) GetValues() apiserver0.Values {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValues")
	ret0, _ := ret[0].(apiserver0.Values)
	return ret0
}

// GetValues indicates an expected call of GetValues.
func (mr *MockInterfaceMockRecorder) GetValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValues", reflect.TypeOf((*MockInterface)(nil).GetValues))
}

// ScrapeConfigs mocks base method.
func (m *MockInterface) ScrapeConfigs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScrapeConfigs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScrapeConfigs indicates an expected call of ScrapeConfigs.
func (mr *MockInterfaceMockRecorder) ScrapeConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScrapeConfigs", reflect.TypeOf((*MockInterface)(nil).ScrapeConfigs))
}

// SetAutoscalingAPIServerResources mocks base method.
func (m *MockInterface) SetAutoscalingAPIServerResources(arg0 v1.ResourceRequirements) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAutoscalingAPIServerResources", arg0)
}

// SetAutoscalingAPIServerResources indicates an expected call of SetAutoscalingAPIServerResources.
func (mr *MockInterfaceMockRecorder) SetAutoscalingAPIServerResources(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAutoscalingAPIServerResources", reflect.TypeOf((*MockInterface)(nil).SetAutoscalingAPIServerResources), arg0)
}

// SetAutoscalingReplicas mocks base method.
func (m *MockInterface) SetAutoscalingReplicas(arg0 *int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAutoscalingReplicas", arg0)
}

// SetAutoscalingReplicas indicates an expected call of SetAutoscalingReplicas.
func (mr *MockInterfaceMockRecorder) SetAutoscalingReplicas(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAutoscalingReplicas", reflect.TypeOf((*MockInterface)(nil).SetAutoscalingReplicas), arg0)
}

// SetETCDEncryptionConfig mocks base method.
func (m *MockInterface) SetETCDEncryptionConfig(arg0 apiserver.ETCDEncryptionConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetETCDEncryptionConfig", arg0)
}

// SetETCDEncryptionConfig indicates an expected call of SetETCDEncryptionConfig.
func (mr *MockInterfaceMockRecorder) SetETCDEncryptionConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetETCDEncryptionConfig", reflect.TypeOf((*MockInterface)(nil).SetETCDEncryptionConfig), arg0)
}

// SetExternalHostname mocks base method.
func (m *MockInterface) SetExternalHostname(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetExternalHostname", arg0)
}

// SetExternalHostname indicates an expected call of SetExternalHostname.
func (mr *MockInterfaceMockRecorder) SetExternalHostname(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExternalHostname", reflect.TypeOf((*MockInterface)(nil).SetExternalHostname), arg0)
}

// SetExternalServer mocks base method.
func (m *MockInterface) SetExternalServer(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetExternalServer", arg0)
}

// SetExternalServer indicates an expected call of SetExternalServer.
func (mr *MockInterfaceMockRecorder) SetExternalServer(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExternalServer", reflect.TypeOf((*MockInterface)(nil).SetExternalServer), arg0)
}

// SetNodeNetworkCIDR mocks base method.
func (m *MockInterface) SetNodeNetworkCIDR(arg0 *string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNodeNetworkCIDR", arg0)
}

// SetNodeNetworkCIDR indicates an expected call of SetNodeNetworkCIDR.
func (mr *MockInterfaceMockRecorder) SetNodeNetworkCIDR(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeNetworkCIDR", reflect.TypeOf((*MockInterface)(nil).SetNodeNetworkCIDR), arg0)
}

// SetSNIConfig mocks base method.
func (m *MockInterface) SetSNIConfig(arg0 apiserver0.SNIConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSNIConfig", arg0)
}

// SetSNIConfig indicates an expected call of SetSNIConfig.
func (mr *MockInterfaceMockRecorder) SetSNIConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSNIConfig", reflect.TypeOf((*MockInterface)(nil).SetSNIConfig), arg0)
}

// SetServerCertificateConfig mocks base method.
func (m *MockInterface) SetServerCertificateConfig(arg0 apiserver0.ServerCertificateConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetServerCertificateConfig", arg0)
}

// SetServerCertificateConfig indicates an expected call of SetServerCertificateConfig.
func (mr *MockInterfaceMockRecorder) SetServerCertificateConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetServerCertificateConfig", reflect.TypeOf((*MockInterface)(nil).SetServerCertificateConfig), arg0)
}

// SetServiceAccountConfig mocks base method.
func (m *MockInterface) SetServiceAccountConfig(arg0 apiserver0.ServiceAccountConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetServiceAccountConfig", arg0)
}

// SetServiceAccountConfig indicates an expected call of SetServiceAccountConfig.
func (mr *MockInterfaceMockRecorder) SetServiceAccountConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetServiceAccountConfig", reflect.TypeOf((*MockInterface)(nil).SetServiceAccountConfig), arg0)
}

// Wait mocks base method.
func (m *MockInterface) Wait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockInterfaceMockRecorder) Wait(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockInterface)(nil).Wait), arg0)
}

// WaitCleanup mocks base method.
func (m *MockInterface) WaitCleanup(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockInterfaceMockRecorder) WaitCleanup(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockInterface)(nil).WaitCleanup), arg0)
}