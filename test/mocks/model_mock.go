// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces/model.go

// Package mocks is a generated GoMock package.
package mocks

import (
	interfaces "linweb/interfaces"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIModel is a mock of IModel interface.
type MockIModel struct {
	ctrl     *gomock.Controller
	recorder *MockIModelMockRecorder
}

// MockIModelMockRecorder is the mock recorder for MockIModel.
type MockIModelMockRecorder struct {
	mock *MockIModel
}

// NewMockIModel creates a new mock instance.
func NewMockIModel(ctrl *gomock.Controller) *MockIModel {
	mock := &MockIModel{ctrl: ctrl}
	mock.recorder = &MockIModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIModel) EXPECT() *MockIModelMockRecorder {
	return m.recorder
}

// MapFromByFieldName mocks base method.
func (m *MockIModel) MapFromByFieldName(src interface{}) interfaces.IModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapFromByFieldName", src)
	ret0, _ := ret[0].(interfaces.IModel)
	return ret0
}

// MapFromByFieldName indicates an expected call of MapFromByFieldName.
func (mr *MockIModelMockRecorder) MapFromByFieldName(src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapFromByFieldName", reflect.TypeOf((*MockIModel)(nil).MapFromByFieldName), src)
}

// MapFromByFieldTag mocks base method.
func (m *MockIModel) MapFromByFieldTag(src interface{}) interfaces.IModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapFromByFieldTag", src)
	ret0, _ := ret[0].(interfaces.IModel)
	return ret0
}

// MapFromByFieldTag indicates an expected call of MapFromByFieldTag.
func (mr *MockIModelMockRecorder) MapFromByFieldTag(src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapFromByFieldTag", reflect.TypeOf((*MockIModel)(nil).MapFromByFieldTag), src)
}

// MapToByFieldName mocks base method.
func (m *MockIModel) MapToByFieldName(dest interface{}) interfaces.IModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapToByFieldName", dest)
	ret0, _ := ret[0].(interfaces.IModel)
	return ret0
}

// MapToByFieldName indicates an expected call of MapToByFieldName.
func (mr *MockIModelMockRecorder) MapToByFieldName(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapToByFieldName", reflect.TypeOf((*MockIModel)(nil).MapToByFieldName), dest)
}

// MapToByFieldTag mocks base method.
func (m *MockIModel) MapToByFieldTag(dest interface{}) interfaces.IModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapToByFieldTag", dest)
	ret0, _ := ret[0].(interfaces.IModel)
	return ret0
}

// MapToByFieldTag indicates an expected call of MapToByFieldTag.
func (mr *MockIModelMockRecorder) MapToByFieldTag(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapToByFieldTag", reflect.TypeOf((*MockIModel)(nil).MapToByFieldTag), dest)
}

// ModelError mocks base method.
func (m *MockIModel) ModelError() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelError")
	ret0, _ := ret[0].(error)
	return ret0
}

// ModelError indicates an expected call of ModelError.
func (mr *MockIModelMockRecorder) ModelError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelError", reflect.TypeOf((*MockIModel)(nil).ModelError))
}

// New mocks base method.
func (m *MockIModel) New(model interface{}) interfaces.IModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", model)
	ret0, _ := ret[0].(interfaces.IModel)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockIModelMockRecorder) New(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockIModel)(nil).New), model)
}

// Validate mocks base method.
func (m *MockIModel) Validate() interfaces.IModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(interfaces.IModel)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockIModelMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockIModel)(nil).Validate))
}
