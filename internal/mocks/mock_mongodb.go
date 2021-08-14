// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/mongo-db/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/sakshi29/getir-app/internal/model"
)

// MockMongoDBRepo is a mock of MongoDBRepo interface.
type MockMongoDBRepo struct {
	ctrl     *gomock.Controller
	recorder *MockMongoDBRepoMockRecorder
}

// MockMongoDBRepoMockRecorder is the mock recorder for MockMongoDBRepo.
type MockMongoDBRepoMockRecorder struct {
	mock *MockMongoDBRepo
}

// NewMockMongoDBRepo creates a new mock instance.
func NewMockMongoDBRepo(ctrl *gomock.Controller) *MockMongoDBRepo {
	mock := &MockMongoDBRepo{ctrl: ctrl}
	mock.recorder = &MockMongoDBRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMongoDBRepo) EXPECT() *MockMongoDBRepoMockRecorder {
	return m.recorder
}

// FetchDocuments mocks base method.
func (m *MockMongoDBRepo) FetchDocuments(ctx context.Context, record model.DocumentFilters) ([]model.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDocuments", ctx, record)
	ret0, _ := ret[0].([]model.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchDocuments indicates an expected call of FetchDocuments.
func (mr *MockMongoDBRepoMockRecorder) FetchDocuments(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDocuments", reflect.TypeOf((*MockMongoDBRepo)(nil).FetchDocuments), ctx, record)
}
