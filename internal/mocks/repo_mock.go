// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-track-api/internal/repo (interfaces: TrackRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	utils "github.com/ozonva/ova-track-api/internal/utils"
)

// MockTrackRepo is a mock of TrackRepo interface.
type MockTrackRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTrackRepoMockRecorder
}

// MockTrackRepoMockRecorder is the mock recorder for MockTrackRepo.
type MockTrackRepoMockRecorder struct {
	mock *MockTrackRepo
}

// NewMockTrackRepo creates a new mock instance.
func NewMockTrackRepo(ctrl *gomock.Controller) *MockTrackRepo {
	mock := &MockTrackRepo{ctrl: ctrl}
	mock.recorder = &MockTrackRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackRepo) EXPECT() *MockTrackRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTrackRepo) Add(arg0 []utils.Track) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockTrackRepoMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTrackRepo)(nil).Add), arg0)
}

// Describe mocks base method.
func (m *MockTrackRepo) Describe(arg0 uint64) (*utils.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Describe", arg0)
	ret0, _ := ret[0].(*utils.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Describe indicates an expected call of Describe.
func (mr *MockTrackRepoMockRecorder) Describe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockTrackRepo)(nil).Describe), arg0)
}

// List mocks base method.
func (m *MockTrackRepo) List(arg0, arg1 uint64) ([]utils.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]utils.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTrackRepoMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTrackRepo)(nil).List), arg0, arg1)
}
