// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/iceCream/src/service (interfaces: IceCreamOperations)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/iceCream/src/model"
)

// Mock of IceCreamOperations interface
type MockIceCreamOperations struct {
	ctrl     *gomock.Controller
	recorder *_MockIceCreamOperationsRecorder
}

// Recorder for MockIceCreamOperations (not exported)
type _MockIceCreamOperationsRecorder struct {
	mock *MockIceCreamOperations
}

func NewMockIceCreamOperations(ctrl *gomock.Controller) *MockIceCreamOperations {
	mock := &MockIceCreamOperations{ctrl: ctrl}
	mock.recorder = &_MockIceCreamOperationsRecorder{mock}
	return mock
}

func (_m *MockIceCreamOperations) EXPECT() *_MockIceCreamOperationsRecorder {
	return _m.recorder
}

func (_m *MockIceCreamOperations) AddIceCreamObj(_param0 model.IceCream) {
	_m.ctrl.Call(_m, "AddIceCreamObj", _param0)
}

func (_mr *_MockIceCreamOperationsRecorder) AddIceCreamObj(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddIceCreamObj", arg0)
}

func (_m *MockIceCreamOperations) GetAllIceCreams() []model.IceCream {
	ret := _m.ctrl.Call(_m, "GetAllIceCreams")
	ret0, _ := ret[0].([]model.IceCream)
	return ret0
}

func (_mr *_MockIceCreamOperationsRecorder) GetAllIceCreams() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllIceCreams")
}

func (_m *MockIceCreamOperations) GetIceCreamByProductID(_param0 string) {
	_m.ctrl.Call(_m, "GetIceCreamByProductID", _param0)
}

func (_mr *_MockIceCreamOperationsRecorder) GetIceCreamByProductID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetIceCreamByProductID", arg0)
}