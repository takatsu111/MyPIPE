// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/repository/PlayListRepository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "MyPIPE/domain/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPlayListRepository is a mock of PlayListRepository interface
type MockPlayListRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPlayListRepositoryMockRecorder
}

// MockPlayListRepositoryMockRecorder is the mock recorder for MockPlayListRepository
type MockPlayListRepositoryMockRecorder struct {
	mock *MockPlayListRepository
}

// NewMockPlayListRepository creates a new mock instance
func NewMockPlayListRepository(ctrl *gomock.Controller) *MockPlayListRepository {
	mock := &MockPlayListRepository{ctrl: ctrl}
	mock.recorder = &MockPlayListRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlayListRepository) EXPECT() *MockPlayListRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method
func (m *MockPlayListRepository) FindByID(playListID model.PlayListID) (*model.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", playListID)
	ret0, _ := ret[0].(*model.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockPlayListRepositoryMockRecorder) FindByID(playListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPlayListRepository)(nil).FindByID), playListID)
}

// FindByName mocks base method
func (m *MockPlayListRepository) FindByName(playListName model.PlayListName) ([]model.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", playListName)
	ret0, _ := ret[0].([]model.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName
func (mr *MockPlayListRepositoryMockRecorder) FindByName(playListName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockPlayListRepository)(nil).FindByName), playListName)
}

// FindByUserID mocks base method
func (m *MockPlayListRepository) FindByUserID(playListUserID model.UserID) ([]model.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", playListUserID)
	ret0, _ := ret[0].([]model.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID
func (mr *MockPlayListRepositoryMockRecorder) FindByUserID(playListUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockPlayListRepository)(nil).FindByUserID), playListUserID)
}

// FindByIDAndUserID mocks base method
func (m *MockPlayListRepository) FindByIDAndUserID(playListID model.PlayListID, userId model.UserID) (*model.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDAndUserID", playListID, userId)
	ret0, _ := ret[0].(*model.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDAndUserID indicates an expected call of FindByIDAndUserID
func (mr *MockPlayListRepositoryMockRecorder) FindByIDAndUserID(playListID, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDAndUserID", reflect.TypeOf((*MockPlayListRepository)(nil).FindByIDAndUserID), playListID, userId)
}

// FindByUserIDAndName mocks base method
func (m *MockPlayListRepository) FindByUserIDAndName(playListUserID model.UserID, playListName model.PlayListName) ([]model.PlayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserIDAndName", playListUserID, playListName)
	ret0, _ := ret[0].([]model.PlayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserIDAndName indicates an expected call of FindByUserIDAndName
func (mr *MockPlayListRepositoryMockRecorder) FindByUserIDAndName(playListUserID, playListName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserIDAndName", reflect.TypeOf((*MockPlayListRepository)(nil).FindByUserIDAndName), playListUserID, playListName)
}

// Save mocks base method
func (m *MockPlayListRepository) Save(playList *model.PlayList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", playList)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockPlayListRepositoryMockRecorder) Save(playList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPlayListRepository)(nil).Save), playList)
}

// Remove mocks base method
func (m *MockPlayListRepository) Remove(userId model.UserID, playListId model.PlayListID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", userId, playListId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockPlayListRepositoryMockRecorder) Remove(userId, playListId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockPlayListRepository)(nil).Remove), userId, playListId)
}
