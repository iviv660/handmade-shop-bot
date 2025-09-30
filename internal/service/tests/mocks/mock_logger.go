package mocks

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func NewMockLogger(t *testing.T) *MockLogger {
	m := &MockLogger{}
	t.Cleanup(func() {
		m.AssertExpectations(t)
	})
	return m
}

func (m *MockLogger) Info(msg string, fields map[string]any) {
	m.Called(msg, fields)
}

func (m *MockLogger) Error(msg string, fields map[string]any) {
	m.Called(msg, fields)
}
