package testutils

import (
	"testing"
)

func testNewMockTCreatesEmptyInstance(t *testing.T) {
	m := NewMockT()

	if len(m.GetErrorf()) != 0 {
		t.Errorf("GetErrorf() = %v, want empty", m.GetErrorf())
	}
	if len(m.GetFatalf()) != 0 {
		t.Errorf("GetFatalf() = %v, want empty", m.GetFatalf())
	}
	if m.GetHelperCalls() != 0 {
		t.Errorf("GetHelperCalls() = %d, want 0", m.GetHelperCalls())
	}
}

func testHelperIncrementsCounter(t *testing.T) {
	m := NewMockT()

	m.Helper()
	if m.GetHelperCalls() != 1 {
		t.Errorf("GetHelperCalls() = %d, want 1", m.GetHelperCalls())
	}

	m.Helper()
	m.Helper()
	if m.GetHelperCalls() != 3 {
		t.Errorf("GetHelperCalls() = %d, want 3", m.GetHelperCalls())
	}
}

func testErrorfCapturesMessages(t *testing.T) {
	m := NewMockT()

	m.Errorf("error %d", 1)
	m.Errorf("error %s", "two")

	errors := m.GetErrorf()
	if len(errors) != 2 {
		t.Fatalf("len(GetErrorf()) = %d, want 2", len(errors))
	}
	if errors[0] != "error 1" {
		t.Errorf("GetErrorf()[0] = %q, want %q", errors[0], "error 1")
	}
	if errors[1] != "error two" {
		t.Errorf("GetErrorf()[1] = %q, want %q", errors[1], "error two")
	}
}

func testFatalfCapturesMessages(t *testing.T) {
	m := NewMockT()

	m.Fatalf("fatal %d", 1)
	m.Fatalf("fatal %s", "two")

	fatalMessages := m.GetFatalf()
	if len(fatalMessages) != 2 {
		t.Fatalf("len(GetFatalf()) = %d, want 2", len(fatalMessages))
	}
	if fatalMessages[0] != "fatal 1" {
		t.Errorf("GetFatalf()[0] = %q, want %q", fatalMessages[0], "fatal 1")
	}
	if fatalMessages[1] != "fatal two" {
		t.Errorf("GetFatalf()[1] = %q, want %q", fatalMessages[1], "fatal two")
	}
}

func testResetClearsAllData(t *testing.T) {
	m := NewMockT()

	// Add some data
	m.Helper()
	m.Helper()
	m.Errorf("error")
	m.Fatalf("fatal")

	// Verify data exists
	if m.GetHelperCalls() != 2 {
		t.Errorf("GetHelperCalls() before reset = %d, want 2", m.GetHelperCalls())
	}
	if len(m.GetErrorf()) != 1 {
		t.Errorf("len(GetErrorf()) before reset = %d, want 1", len(m.GetErrorf()))
	}
	if len(m.GetFatalf()) != 1 {
		t.Errorf("len(GetFatalf()) before reset = %d, want 1", len(m.GetFatalf()))
	}

	// Reset
	m.Reset()

	// Verify all cleared
	if m.GetHelperCalls() != 0 {
		t.Errorf("GetHelperCalls() after reset = %d, want 0", m.GetHelperCalls())
	}
	if len(m.GetErrorf()) != 0 {
		t.Errorf("len(GetErrorf()) after reset = %d, want 0", len(m.GetErrorf()))
	}
	if len(m.GetFatalf()) != 0 {
		t.Errorf("len(GetFatalf()) after reset = %d, want 0", len(m.GetFatalf()))
	}
}

func testConcurrentSafety(t *testing.T) {
	// MockT is designed for single-threaded test usage
	// This test documents that it's NOT safe for concurrent use
	m := NewMockT()

	// Sequential calls work fine
	for i := 0; i < 10; i++ {
		m.Errorf("error %d", i)
	}

	if len(m.GetErrorf()) != 10 {
		t.Errorf("len(GetErrorf()) = %d, want 10", len(m.GetErrorf()))
	}
}

func TestMockT(t *testing.T) {
	t.Run("NewMockT creates empty instance", testNewMockTCreatesEmptyInstance)
	t.Run("Helper increments counter", testHelperIncrementsCounter)
	t.Run("Errorf captures messages", testErrorfCapturesMessages)
	t.Run("Fatalf captures messages", testFatalfCapturesMessages)
	t.Run("Reset clears all data", testResetClearsAllData)
	t.Run("concurrent safety", testConcurrentSafety)
}
