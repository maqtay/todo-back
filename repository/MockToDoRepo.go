package repository

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockRepo struct {
	controller *gomock.Controller
	recorder *MockRepoMockRec
}

type MockRepoMockRec struct {
	mock *MockRepo
}

func NewMockRepo(controller *gomock.Controller) *MockRepo  {
	mock := &MockRepo{controller: controller}
	mock.recorder = &MockRepoMockRec{mock}
	return mock
}

func (m *MockRepo) EXPECT() *MockRepoMockRec {
	return m.recorder
}

func (m *MockRepo) GetAll() interface{} {
	m.controller.T.Helper()
	ret := m.controller.Call(m, "GetAll")
	actual, _ := ret[0].(interface{})
	return actual
}

func (recorder *MockRepoMockRec) GetALl() *gomock.Call {
	recorder.mock.controller.T.Helper()
	return recorder.mock.controller.RecordCallWithMethodType(recorder.mock, "GetAll", reflect.TypeOf((*MockRepoMockRec)(nil).GetALl()))
}

func (m *MockRepo) Add(note string) interface{} {
	m.controller.T.Helper()
	ret := m.controller.Call(m, "Add", note)
	actual := ret[0].(interface{})
	return actual
}

func (recorder *MockRepoMockRec) Add(note string) *gomock.Call  {
	recorder.mock.controller.T.Helper()
	return recorder.mock.controller.RecordCallWithMethodType(recorder.mock, "Add", reflect.TypeOf((*MockRepo)(nil).Add), note)
}

func (m *MockRepo) Delete(id int) interface{}  {
	m.controller.T.Helper()
	ret := m.controller.Call(m, "Delete", id)
	actual := ret[0].(interface{})
	return actual
}

func (recorder *MockRepoMockRec) Delete(id int) *gomock.Call  {
	recorder.mock.controller.T.Helper()
	return recorder.mock.controller.RecordCallWithMethodType(recorder.mock, "Delete", reflect.TypeOf((*MockRepo)(nil).Delete), id)
}
