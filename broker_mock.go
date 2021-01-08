// Code generated by MockGen. DO NOT EDIT.
// Source: broker.go

// Package main is a generated GoMock package.
package main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockstorer is a mock of storer interface
type Mockstorer struct {
	ctrl     *gomock.Controller
	recorder *MockstorerMockRecorder
}

// MockstorerMockRecorder is the mock recorder for Mockstorer
type MockstorerMockRecorder struct {
	mock *Mockstorer
}

// NewMockstorer creates a new mock instance
func NewMockstorer(ctrl *gomock.Controller) *Mockstorer {
	mock := &Mockstorer{ctrl: ctrl}
	mock.recorder = &MockstorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockstorer) EXPECT() *MockstorerMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *Mockstorer) Insert(topic string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", topic, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockstorerMockRecorder) Insert(topic, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*Mockstorer)(nil).Insert), topic, value)
}

// Close mocks base method
func (m *Mockstorer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockstorerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Mockstorer)(nil).Close))
}
