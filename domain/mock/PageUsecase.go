// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/page.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/unknowntpo/page/domain"
)

// MockPageUsecase is a mock of PageUsecase interface.
type MockPageUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPageUsecaseMockRecorder
}

// MockPageUsecaseMockRecorder is the mock recorder for MockPageUsecase.
type MockPageUsecaseMockRecorder struct {
	mock *MockPageUsecase
}

// NewMockPageUsecase creates a new mock instance.
func NewMockPageUsecase(ctrl *gomock.Controller) *MockPageUsecase {
	mock := &MockPageUsecase{ctrl: ctrl}
	mock.recorder = &MockPageUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPageUsecase) EXPECT() *MockPageUsecaseMockRecorder {
	return m.recorder
}

// GetHead mocks base method.
func (m *MockPageUsecase) GetHead(ctx context.Context, userID int64, listKey domain.ListKey) (domain.PageKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead", ctx, userID, listKey)
	ret0, _ := ret[0].(domain.PageKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead.
func (mr *MockPageUsecaseMockRecorder) GetHead(ctx, userID, listKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockPageUsecase)(nil).GetHead), ctx, userID, listKey)
}

// GetPage mocks base method.
func (m *MockPageUsecase) GetPage(ctx context.Context, userID int64, listKey domain.ListKey, pageKey domain.PageKey) (domain.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPage", ctx, userID, listKey, pageKey)
	ret0, _ := ret[0].(domain.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPage indicates an expected call of GetPage.
func (mr *MockPageUsecaseMockRecorder) GetPage(ctx, userID, listKey, pageKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPage", reflect.TypeOf((*MockPageUsecase)(nil).GetPage), ctx, userID, listKey, pageKey)
}

// NewList mocks base method.
func (m *MockPageUsecase) NewList(ctx context.Context, userID int64, listKey domain.ListKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewList", ctx, userID, listKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewList indicates an expected call of NewList.
func (mr *MockPageUsecaseMockRecorder) NewList(ctx, userID, listKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewList", reflect.TypeOf((*MockPageUsecase)(nil).NewList), ctx, userID, listKey)
}

// SetPage mocks base method.
func (m *MockPageUsecase) SetPage(ctx context.Context, userID int64, listKey domain.ListKey, page domain.Page) (domain.PageKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPage", ctx, userID, listKey, page)
	ret0, _ := ret[0].(domain.PageKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetPage indicates an expected call of SetPage.
func (mr *MockPageUsecaseMockRecorder) SetPage(ctx, userID, listKey, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPage", reflect.TypeOf((*MockPageUsecase)(nil).SetPage), ctx, userID, listKey, page)
}

// MockPageRepo is a mock of PageRepo interface.
type MockPageRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPageRepoMockRecorder
}

// MockPageRepoMockRecorder is the mock recorder for MockPageRepo.
type MockPageRepoMockRecorder struct {
	mock *MockPageRepo
}

// NewMockPageRepo creates a new mock instance.
func NewMockPageRepo(ctrl *gomock.Controller) *MockPageRepo {
	mock := &MockPageRepo{ctrl: ctrl}
	mock.recorder = &MockPageRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPageRepo) EXPECT() *MockPageRepoMockRecorder {
	return m.recorder
}

// GetHead mocks base method.
func (m *MockPageRepo) GetHead(ctx context.Context, userID int64, listKey domain.ListKey) (domain.PageKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead", ctx, userID, listKey)
	ret0, _ := ret[0].(domain.PageKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead.
func (mr *MockPageRepoMockRecorder) GetHead(ctx, userID, listKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockPageRepo)(nil).GetHead), ctx, userID, listKey)
}

// GetPage mocks base method.
func (m *MockPageRepo) GetPage(ctx context.Context, uesrID int64, listKey domain.ListKey, pageKey domain.PageKey) (domain.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPage", ctx, uesrID, listKey, pageKey)
	ret0, _ := ret[0].(domain.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPage indicates an expected call of GetPage.
func (mr *MockPageRepoMockRecorder) GetPage(ctx, uesrID, listKey, pageKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPage", reflect.TypeOf((*MockPageRepo)(nil).GetPage), ctx, uesrID, listKey, pageKey)
}

// NewList mocks base method.
func (m *MockPageRepo) NewList(ctx context.Context, userID int64, listKey domain.ListKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewList", ctx, userID, listKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewList indicates an expected call of NewList.
func (mr *MockPageRepoMockRecorder) NewList(ctx, userID, listKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewList", reflect.TypeOf((*MockPageRepo)(nil).NewList), ctx, userID, listKey)
}

// SetPage mocks base method.
func (m *MockPageRepo) SetPage(ctx context.Context, userID int64, listkey domain.ListKey, page domain.Page) (domain.PageKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPage", ctx, userID, listkey, page)
	ret0, _ := ret[0].(domain.PageKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetPage indicates an expected call of SetPage.
func (mr *MockPageRepoMockRecorder) SetPage(ctx, userID, listkey, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPage", reflect.TypeOf((*MockPageRepo)(nil).SetPage), ctx, userID, listkey, page)
}
