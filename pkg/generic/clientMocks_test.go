// Code generated by MockGen. DO NOT EDIT.
// Source: ./embeddedClient.go

// Package generic is a generated GoMock package.
package generic

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
)

// MockembeddedClient is a mock of embeddedClient interface.
type MockembeddedClient struct {
	ctrl     *gomock.Controller
	recorder *MockembeddedClientMockRecorder
}

// MockembeddedClientMockRecorder is the mock recorder for MockembeddedClient.
type MockembeddedClientMockRecorder struct {
	mock *MockembeddedClient
}

// NewMockembeddedClient creates a new mock instance.
func NewMockembeddedClient(ctrl *gomock.Controller) *MockembeddedClient {
	mock := &MockembeddedClient{ctrl: ctrl}
	mock.recorder = &MockembeddedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockembeddedClient) EXPECT() *MockembeddedClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockembeddedClient) Create(ctx context.Context, namespace string, obj, result runtime.Object, opts v1.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, namespace, obj, result, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockembeddedClientMockRecorder) Create(ctx, namespace, obj, result, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockembeddedClient)(nil).Create), ctx, namespace, obj, result, opts)
}

// Delete mocks base method.
func (m *MockembeddedClient) Delete(ctx context.Context, namespace, name string, opts v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, namespace, name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockembeddedClientMockRecorder) Delete(ctx, namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockembeddedClient)(nil).Delete), ctx, namespace, name, opts)
}

// Get mocks base method.
func (m *MockembeddedClient) Get(ctx context.Context, namespace, name string, result runtime.Object, options v1.GetOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, namespace, name, result, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockembeddedClientMockRecorder) Get(ctx, namespace, name, result, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockembeddedClient)(nil).Get), ctx, namespace, name, result, options)
}

// List mocks base method.
func (m *MockembeddedClient) List(ctx context.Context, namespace string, result runtime.Object, opts v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, namespace, result, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockembeddedClientMockRecorder) List(ctx, namespace, result, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockembeddedClient)(nil).List), ctx, namespace, result, opts)
}

// Patch mocks base method.
func (m *MockembeddedClient) Patch(ctx context.Context, namespace, name string, pt types.PatchType, data []byte, result runtime.Object, opts v1.PatchOptions, subresources ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, namespace, name, pt, data, result, opts}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockembeddedClientMockRecorder) Patch(ctx, namespace, name, pt, data, result, opts interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, namespace, name, pt, data, result, opts}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockembeddedClient)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockembeddedClient) Update(ctx context.Context, namespace string, obj, result runtime.Object, opts v1.UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, namespace, obj, result, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockembeddedClientMockRecorder) Update(ctx, namespace, obj, result, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockembeddedClient)(nil).Update), ctx, namespace, obj, result, opts)
}

// UpdateStatus mocks base method.
func (m *MockembeddedClient) UpdateStatus(ctx context.Context, namespace string, obj, result runtime.Object, opts v1.UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, namespace, obj, result, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockembeddedClientMockRecorder) UpdateStatus(ctx, namespace, obj, result, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockembeddedClient)(nil).UpdateStatus), ctx, namespace, obj, result, opts)
}

// Watch mocks base method.
func (m *MockembeddedClient) Watch(ctx context.Context, namespace string, opts v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", ctx, namespace, opts)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockembeddedClientMockRecorder) Watch(ctx, namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockembeddedClient)(nil).Watch), ctx, namespace, opts)
}
