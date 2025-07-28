// Package testutils provides testing utilities for protomcp.org projects.
//
// # T Interface
//
// The T interface is a minimal subset of testing.TB that provides the core
// methods needed for test assertions and helpers:
//
//	type T interface {
//	    Helper()
//	    Errorf(format string, args ...any)
//	    Fatalf(format string, args ...any)
//	}
//
// This interface is satisfied by *testing.T, making it easy to use in tests
// while also allowing for mock implementations.
//
// # MockT
//
// MockT is a test implementation of the T interface, useful for testing
// test helpers and assertion libraries:
//
//	m := testutils.NewMockT()
//	// use m in place of *testing.T
//	// then inspect what happened:
//	errors := m.GetErrorf()     // get all Errorf calls
//	fatalErrors := m.GetFatalf()     // get all Fatalf calls
//	helpers := m.GetHelperCalls() // count Helper() calls
//
// MockT provides methods to inspect test behaviour:
//   - GetErrorf() []string - returns all error messages from Errorf calls
//   - GetFatalf() []string - returns all fatal messages from Fatalf calls
//   - GetHelperCalls() int - returns the number of times Helper() was called
//   - Reset() - clears all recorded data for reuse
//
// Example usage:
//
//	func TestMyHelper(t *testing.T) {
//	    m := testutils.NewMockT()
//	    MyHelper(m, someValue)
//
//	    if len(m.GetErrorf()) > 0 {
//	        t.Errorf("unexpected errors: %v", m.GetErrorf())
//	    }
//	}
package testutils
