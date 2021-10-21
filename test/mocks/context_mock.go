// Code generated by MockGen. DO NOT EDIT.
// Source: ../interfaces/context.go

// Package mocks is a generated GoMock package.
package mocks

import (
	interfaces "github.com/Codexiaoyi/linweb/interfaces"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIContext is a mock of IContext interface.
type MockIContext struct {
	ctrl     *gomock.Controller
	recorder *MockIContextMockRecorder
}

// MockIContextMockRecorder is the mock recorder for MockIContext.
type MockIContextMockRecorder struct {
	mock *MockIContext
}

// NewMockIContext creates a new mock instance.
func NewMockIContext(ctrl *gomock.Controller) *MockIContext {
	mock := &MockIContext{ctrl: ctrl}
	mock.recorder = &MockIContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIContext) EXPECT() *MockIContextMockRecorder {
	return m.recorder
}

// Middleware mocks base method.
func (m *MockIContext) Middleware() interfaces.IMiddleware {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Middleware")
	ret0, _ := ret[0].(interfaces.IMiddleware)
	return ret0
}

// Middleware indicates an expected call of Middleware.
func (mr *MockIContextMockRecorder) Middleware() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Middleware", reflect.TypeOf((*MockIContext)(nil).Middleware))
}

// New mocks base method.
func (m *MockIContext) New() interfaces.IContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New")
	ret0, _ := ret[0].(interfaces.IContext)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockIContextMockRecorder) New() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockIContext)(nil).New))
}

// Next mocks base method.
func (m *MockIContext) Next() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Next")
}

// Next indicates an expected call of Next.
func (mr *MockIContextMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIContext)(nil).Next))
}

// Request mocks base method.
func (m *MockIContext) Request() interfaces.IRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(interfaces.IRequest)
	return ret0
}

// Request indicates an expected call of Request.
func (mr *MockIContextMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockIContext)(nil).Request))
}

// Reset mocks base method.
func (m *MockIContext) Reset(arg0 http.ResponseWriter, arg1 *http.Request, arg2 interfaces.IMiddleware) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", arg0, arg1, arg2)
}

// Reset indicates an expected call of Reset.
func (mr *MockIContextMockRecorder) Reset(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockIContext)(nil).Reset), arg0, arg1, arg2)
}

// Response mocks base method.
func (m *MockIContext) Response() interfaces.IResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Response")
	ret0, _ := ret[0].(interfaces.IResponse)
	return ret0
}

// Response indicates an expected call of Response.
func (mr *MockIContextMockRecorder) Response() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockIContext)(nil).Response))
}

// MockIRequest is a mock of IRequest interface.
type MockIRequest struct {
	ctrl     *gomock.Controller
	recorder *MockIRequestMockRecorder
}

// MockIRequestMockRecorder is the mock recorder for MockIRequest.
type MockIRequestMockRecorder struct {
	mock *MockIRequest
}

// NewMockIRequest creates a new mock instance.
func NewMockIRequest(ctrl *gomock.Controller) *MockIRequest {
	mock := &MockIRequest{ctrl: ctrl}
	mock.recorder = &MockIRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRequest) EXPECT() *MockIRequestMockRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockIRequest) Body() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].(string)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockIRequestMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockIRequest)(nil).Body))
}

// Method mocks base method.
func (m *MockIRequest) Method() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(string)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockIRequestMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockIRequest)(nil).Method))
}

// Param mocks base method.
func (m *MockIRequest) Param(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param.
func (mr *MockIRequestMockRecorder) Param(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockIRequest)(nil).Param), arg0)
}

// Path mocks base method.
func (m *MockIRequest) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockIRequestMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockIRequest)(nil).Path))
}

// PostForm mocks base method.
func (m *MockIRequest) PostForm(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostForm", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// PostForm indicates an expected call of PostForm.
func (mr *MockIRequestMockRecorder) PostForm(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockIRequest)(nil).PostForm), key)
}

// Query mocks base method.
func (m *MockIRequest) Query(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockIRequestMockRecorder) Query(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockIRequest)(nil).Query), key)
}

// SetParams mocks base method.
func (m *MockIRequest) SetParams(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParams", arg0)
}

// SetParams indicates an expected call of SetParams.
func (mr *MockIRequestMockRecorder) SetParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParams", reflect.TypeOf((*MockIRequest)(nil).SetParams), arg0)
}

// MockIResponse is a mock of IResponse interface.
type MockIResponse struct {
	ctrl     *gomock.Controller
	recorder *MockIResponseMockRecorder
}

// MockIResponseMockRecorder is the mock recorder for MockIResponse.
type MockIResponseMockRecorder struct {
	mock *MockIResponse
}

// NewMockIResponse creates a new mock instance.
func NewMockIResponse(ctrl *gomock.Controller) *MockIResponse {
	mock := &MockIResponse{ctrl: ctrl}
	mock.recorder = &MockIResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIResponse) EXPECT() *MockIResponseMockRecorder {
	return m.recorder
}

// Data mocks base method.
func (m *MockIResponse) Data(code int, data []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Data", code, data)
}

// Data indicates an expected call of Data.
func (mr *MockIResponseMockRecorder) Data(code, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockIResponse)(nil).Data), code, data)
}

// HTML mocks base method.
func (m *MockIResponse) HTML(code int, html string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HTML", code, html)
}

// HTML indicates an expected call of HTML.
func (mr *MockIResponseMockRecorder) HTML(code, html interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTML", reflect.TypeOf((*MockIResponse)(nil).HTML), code, html)
}

// Header mocks base method.
func (m *MockIResponse) Header(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Header", key, value)
}

// Header indicates an expected call of Header.
func (mr *MockIResponseMockRecorder) Header(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockIResponse)(nil).Header), key, value)
}

// JSON mocks base method.
func (m *MockIResponse) JSON(code int, obj interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JSON", code, obj)
}

// JSON indicates an expected call of JSON.
func (mr *MockIResponseMockRecorder) JSON(code, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockIResponse)(nil).JSON), code, obj)
}

// Status mocks base method.
func (m *MockIResponse) Status(code int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Status", code)
}

// Status indicates an expected call of Status.
func (mr *MockIResponseMockRecorder) Status(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockIResponse)(nil).Status), code)
}

// String mocks base method.
func (m *MockIResponse) String(code int, format string, values ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{code, format}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "String", varargs...)
}

// String indicates an expected call of String.
func (mr *MockIResponseMockRecorder) String(code, format interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{code, format}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockIResponse)(nil).String), varargs...)
}
