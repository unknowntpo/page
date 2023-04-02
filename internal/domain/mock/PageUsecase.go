// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/page.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/unknowntpo/page/internal/domain"
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
func (m *MockPageUsecase) GetHead(ctx context.Context, listKey domain.ListKey) (domain.PageKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead", ctx, listKey)
	ret0, _ := ret[0].(domain.PageKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead.
func (mr *MockPageUsecaseMockRecorder) GetHead(ctx, listKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockPageUsecase)(nil).GetHead), ctx, listKey)
}

// GetPage mocks base method.
func (m *MockPageUsecase) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPage", ctx, pageKey)
	ret0, _ := ret[0].(domain.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPage indicates an expected call of GetPage.
func (mr *MockPageUsecaseMockRecorder) GetPage(ctx, pageKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPage", reflect.TypeOf((*MockPageUsecase)(nil).GetPage), ctx, pageKey)
}

// SetPage mocks base method.
func (m *MockPageUsecase) SetPage(ctx context.Context, page domain.Page) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPage", ctx, page)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPage indicates an expected call of SetPage.
func (mr *MockPageUsecaseMockRecorder) SetPage(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPage", reflect.TypeOf((*MockPageUsecase)(nil).SetPage), ctx, page)
}

// MockPageAPI is a mock of PageAPI interface.
type MockPageAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPageAPIMockRecorder
}

// MockPageAPIMockRecorder is the mock recorder for MockPageAPI.
type MockPageAPIMockRecorder struct {
	mock *MockPageAPI
}

// NewMockPageAPI creates a new mock instance.
func NewMockPageAPI(ctrl *gomock.Controller) *MockPageAPI {
	mock := &MockPageAPI{ctrl: ctrl}
	mock.recorder = &MockPageAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPageAPI) EXPECT() *MockPageAPIMockRecorder {
	return m.recorder
}

// GetHead mocks base method.
func (m *MockPageAPI) GetHead(ctx context.Context, listKey domain.ListKey) (domain.PageKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead", ctx, listKey)
	ret0, _ := ret[0].(domain.PageKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead.
func (mr *MockPageAPIMockRecorder) GetHead(ctx, listKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockPageAPI)(nil).GetHead), ctx, listKey)
}

// GetPage mocks base method.
func (m *MockPageAPI) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPage", ctx, pageKey)
	ret0, _ := ret[0].(domain.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPage indicates an expected call of GetPage.
func (mr *MockPageAPIMockRecorder) GetPage(ctx, pageKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPage", reflect.TypeOf((*MockPageAPI)(nil).GetPage), ctx, pageKey)
}

// SetPage mocks base method.
func (m *MockPageAPI) SetPage(ctx context.Context, page domain.Page) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPage", ctx, page)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPage indicates an expected call of SetPage.
func (mr *MockPageAPIMockRecorder) SetPage(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPage", reflect.TypeOf((*MockPageAPI)(nil).SetPage), ctx, page)
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
func (m *MockPageRepo) GetHead(ctx context.Context, listKey domain.ListKey) (domain.PageKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead", ctx, listKey)
	ret0, _ := ret[0].(domain.PageKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead.
func (mr *MockPageRepoMockRecorder) GetHead(ctx, listKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockPageRepo)(nil).GetHead), ctx, listKey)
}

// GetPage mocks base method.
func (m *MockPageRepo) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPage", ctx, pageKey)
	ret0, _ := ret[0].(domain.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPage indicates an expected call of GetPage.
func (mr *MockPageRepoMockRecorder) GetPage(ctx, pageKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPage", reflect.TypeOf((*MockPageRepo)(nil).GetPage), ctx, pageKey)
}

// SetPage mocks base method.
func (m *MockPageRepo) SetPage(ctx context.Context, listkey domain.ListKey, page domain.Page) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPage", ctx, listkey, page)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPage indicates an expected call of SetPage.
func (mr *MockPageRepoMockRecorder) SetPage(ctx, listkey, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPage", reflect.TypeOf((*MockPageRepo)(nil).SetPage), ctx, listkey, page)
}
