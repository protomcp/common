package testutils

import (
	"fmt"
	"testing"
)

// T is a subset of testing.TB for assertions
type T interface {
	Helper()
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

// Ensure *testing.T satisfies our T interface
var _ T = (*testing.T)(nil)

// Ensure MockT satisfies our T interface
var _ T = (*MockT)(nil)

// MockT implements T interface for testing assertions
type MockT struct {
	errorf []string
	fatalf []string
	helper int
}

// NewMockT creates a new MockT instance
func NewMockT() *MockT {
	return &MockT{}
}

// Helper implements T.Helper by incrementing the call counter
func (m *MockT) Helper() {
	m.helper++
}

// Errorf implements T.Errorf by recording the formatted message
func (m *MockT) Errorf(format string, args ...any) {
	m.errorf = append(m.errorf, fmt.Sprintf(format, args...))
}

// Fatalf implements T.Fatalf by recording the formatted message
func (m *MockT) Fatalf(format string, args ...any) {
	m.fatalf = append(m.fatalf, fmt.Sprintf(format, args...))
}

// GetErrorf returns all error messages
func (m *MockT) GetErrorf() []string {
	return m.errorf
}

// GetFatalf returns all fatal messages
func (m *MockT) GetFatalf() []string {
	return m.fatalf
}

// GetHelperCalls returns the number of Helper() calls
func (m *MockT) GetHelperCalls() int {
	return m.helper
}

// Reset clears all recorded data
func (m *MockT) Reset() {
	m.errorf = nil
	m.fatalf = nil
	m.helper = 0
}
