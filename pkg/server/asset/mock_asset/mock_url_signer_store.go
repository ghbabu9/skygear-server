// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/skygeario/skygear-server/pkg/server/asset (interfaces: URLSignerStore)

package mock_asset

import (
	gomock "github.com/golang/mock/gomock"
	asset "github.com/skygeario/skygear-server/pkg/server/asset"
	io "io"
	reflect "reflect"
)

// MockURLSignerStore is a mock of URLSignerStore interface
type MockURLSignerStore struct {
	ctrl     *gomock.Controller
	recorder *MockURLSignerStoreMockRecorder
}

// MockURLSignerStoreMockRecorder is the mock recorder for MockURLSignerStore
type MockURLSignerStoreMockRecorder struct {
	mock *MockURLSignerStore
}

// NewMockURLSignerStore creates a new mock instance
func NewMockURLSignerStore(ctrl *gomock.Controller) *MockURLSignerStore {
	mock := &MockURLSignerStore{ctrl: ctrl}
	mock.recorder = &MockURLSignerStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockURLSignerStore) EXPECT() *MockURLSignerStoreMockRecorder {
	return _m.recorder
}

// GeneratePostFileRequest mocks base method
func (_m *MockURLSignerStore) GeneratePostFileRequest(_param0 string) (*asset.PostFileRequest, error) {
	ret := _m.ctrl.Call(_m, "GeneratePostFileRequest", _param0)
	ret0, _ := ret[0].(*asset.PostFileRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePostFileRequest indicates an expected call of GeneratePostFileRequest
func (_mr *MockURLSignerStoreMockRecorder) GeneratePostFileRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GeneratePostFileRequest", reflect.TypeOf((*MockURLSignerStore)(nil).GeneratePostFileRequest), arg0)
}

// GetFileReader mocks base method
func (_m *MockURLSignerStore) GetFileReader(_param0 string) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "GetFileReader", _param0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileReader indicates an expected call of GetFileReader
func (_mr *MockURLSignerStoreMockRecorder) GetFileReader(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetFileReader", reflect.TypeOf((*MockURLSignerStore)(nil).GetFileReader), arg0)
}

// IsSignatureRequired mocks base method
func (_m *MockURLSignerStore) IsSignatureRequired() bool {
	ret := _m.ctrl.Call(_m, "IsSignatureRequired")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSignatureRequired indicates an expected call of IsSignatureRequired
func (_mr *MockURLSignerStoreMockRecorder) IsSignatureRequired() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IsSignatureRequired", reflect.TypeOf((*MockURLSignerStore)(nil).IsSignatureRequired))
}

// PutFileReader mocks base method
func (_m *MockURLSignerStore) PutFileReader(_param0 string, _param1 io.Reader, _param2 int64, _param3 string) error {
	ret := _m.ctrl.Call(_m, "PutFileReader", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutFileReader indicates an expected call of PutFileReader
func (_mr *MockURLSignerStoreMockRecorder) PutFileReader(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "PutFileReader", reflect.TypeOf((*MockURLSignerStore)(nil).PutFileReader), arg0, arg1, arg2, arg3)
}

// SignedURL mocks base method
func (_m *MockURLSignerStore) SignedURL(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "SignedURL", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignedURL indicates an expected call of SignedURL
func (_mr *MockURLSignerStoreMockRecorder) SignedURL(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SignedURL", reflect.TypeOf((*MockURLSignerStore)(nil).SignedURL), arg0)
}
