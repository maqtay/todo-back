package services

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockService struct {
	controller *gomock.Controller
	recorder *MockServiceMockRec
}

type MockServiceMockRec struct {
	mock *MockService
}

func NewMockService(controller *gomock.Controller) *MockService  {
	mock := &MockService{controller: controller}
	mock.recorder = &MockServiceMockRec{mock}
	return mock
}

func (m *MockService) EXPECT() *MockServiceMockRec {
	return m.recorder
}

func (m *MockService) GetAll() interface{} {
	m.controller.T.Helper()
	ret := m.controller.Call(m, "GetAll")
	actual, _ := ret[0].(interface{})
	return actual
}

func (recorder *MockServiceMockRec) GetALl() *gomock.Call {
	recorder.mock.controller.T.Helper()
	return recorder.mock.controller.RecordCallWithMethodType(recorder.mock, "GetAll", reflect.TypeOf((*MockServiceMockRec)(nil).GetALl()))
}

func (m *MockService) Add(note string) interface{} {
	m.controller.T.Helper()
	ret := m.controller.Call(m, "Add", note)
	actual := ret[0].(interface{})
	return actual
}

func (recorder *MockServiceMockRec) Add(note string) *gomock.Call  {
	recorder.mock.controller.T.Helper()
	return recorder.mock.controller.RecordCallWithMethodType(recorder.mock, "Add", reflect.TypeOf((*MockService)(nil).Add), note)
}

func (m *MockService) Delete(id int) interface{}  {
	m.controller.T.Helper()
	ret := m.controller.Call(m, "Delete", id)
	actual := ret[0].(interface{})
	return actual
}

func (recorder *MockServiceMockRec) Delete(id int) *gomock.Call  {
	recorder.mock.controller.T.Helper()
	return recorder.mock.controller.RecordCallWithMethodType(recorder.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), id)
}

