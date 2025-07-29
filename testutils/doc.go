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
//
// # Assertion Functions
//
// The package provides a comprehensive set of assertion functions that all
// return bool to indicate success, allowing them to be used in conditional
// logic:
//
//	if !testutils.AssertEqual(t, expected, actual, "config value") {
//	    // handle failure case
//	    return
//	}
//
// Available assertions:
//   - AssertEqual[V any](t T, expected, actual V, name string, args ...any) bool
//   - AssertNotEqual[V any](t T, expected, actual V, name string, args ...any) bool
//   - AssertNil(t T, value any, name string, args ...any) bool
//   - AssertNotNil(t T, value any, name string, args ...any) bool
//   - AssertTrue(t T, value bool, name string, args ...any) bool
//   - AssertFalse(t T, value bool, name string, args ...any) bool
//   - AssertError(t T, err error, name string, args ...any) bool
//   - AssertNoError(t T, err error, name string, args ...any) bool
//   - AssertErrorIs(t T, err, target error, name string, args ...any) bool
//   - AssertContains(t T, str, substr string, name string, args ...any) bool
//   - AssertTypeIs[U any](t T, value any, name string, args ...any) (U, bool)
//
// All assertion functions support formatted names:
//
//	testutils.AssertEqual(t, 42, value, "field %s value", fieldName)
//
// # Helper Functions
//
// IsNil checks if a value is nil using reflection, properly handling
// nil interfaces, nil pointers, and other nil-able types:
//
//	if testutils.IsNil(value) {
//	    // handle nil case
//	}
package testutils
