// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/yuribrito/go/src/github.com/Shopify/sarama/sync_producer.go

// Package mock_sarama is a generated GoMock package.
package mock_sarama

import (
	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSyncProducer is a mock of SyncProducer interface
type MockSyncProducer struct {
	ctrl     *gomock.Controller
	recorder *MockSyncProducerMockRecorder
}

// MockSyncProducerMockRecorder is the mock recorder for MockSyncProducer
type MockSyncProducerMockRecorder struct {
	mock *MockSyncProducer
}

// NewMockSyncProducer creates a new mock instance
func NewMockSyncProducer(ctrl *gomock.Controller) *MockSyncProducer {
	mock := &MockSyncProducer{ctrl: ctrl}
	mock.recorder = &MockSyncProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSyncProducer) EXPECT() *MockSyncProducerMockRecorder {
	return m.recorder
}

// SendMessage mocks base method
func (m *MockSyncProducer) SendMessage(msg *sarama.ProducerMessage) (int32, int64, error) {
	ret := m.ctrl.Call(m, "SendMessage", msg)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockSyncProducerMockRecorder) SendMessage(msg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockSyncProducer)(nil).SendMessage), msg)
}

// SendMessages mocks base method
func (m *MockSyncProducer) SendMessages(msgs []*sarama.ProducerMessage) error {
	ret := m.ctrl.Call(m, "SendMessages", msgs)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessages indicates an expected call of SendMessages
func (mr *MockSyncProducerMockRecorder) SendMessages(msgs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockSyncProducer)(nil).SendMessages), msgs)
}

// Close mocks base method
func (m *MockSyncProducer) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSyncProducerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncProducer)(nil).Close))
}
