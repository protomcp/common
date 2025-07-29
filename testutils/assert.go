package testutils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// AssertEqual fails the test if expected and actual are not equal.
// Returns true if the assertion passed, false otherwise.
// The name parameter can include printf-style formatting.
func AssertEqual[V any](t T, expected, actual V, name string, args ...any) bool {
	t.Helper()
	ok := reflect.DeepEqual(expected, actual)
	if !ok {
		doError(t, name, args, "expected %v, got %v", expected, actual)
	}
	return ok
}

// AssertNotEqual fails the test if expected and actual are equal.
// Returns true if the assertion passed, false otherwise.
func AssertNotEqual[V any](t T, expected, actual V, name string, args ...any) bool {
	t.Helper()
	ok := !reflect.DeepEqual(expected, actual)
	if !ok {
		doError(t, name, args, "expected values to be different, both were %v", expected)
	}
	return ok
}

// AssertNil fails the test if value is not nil.
// Returns true if the assertion passed, false otherwise.
func AssertNil(t T, value any, name string, args ...any) bool {
	t.Helper()
	ok := IsNil(value)
	if !ok {
		doError(t, name, args, "expected nil, got %v", value)
	}
	return ok
}

// AssertNotNil fails the test if value is nil.
// Returns true if the assertion passed, false otherwise.
func AssertNotNil(t T, value any, name string, args ...any) bool {
	t.Helper()
	ok := !IsNil(value)
	if !ok {
		doError(t, name, args, "expected non-nil value")
	}
	return ok
}

// AssertTrue fails the test if value is not true.
// Returns true if the assertion passed, false otherwise.
//
//revive:disable-next-line:flag-parameter
func AssertTrue(t T, value bool, name string, args ...any) bool {
	t.Helper()
	if !value {
		doError(t, name, args, "expected true, got false")
	}
	return value
}

// AssertFalse fails the test if value is not false.
// Returns true if the assertion passed, false otherwise.
//
//revive:disable-next-line:flag-parameter
func AssertFalse(t T, value bool, name string, args ...any) bool {
	t.Helper()
	ok := !value
	if !ok {
		doError(t, name, args, "expected false, got true")
	}
	return ok
}

// AssertErrorIs fails the test if the error does not match the target error.
// Uses errors.Is to check if the error matches.
// Returns true if the assertion passed, false otherwise.
func AssertErrorIs(t T, err, target error, name string, args ...any) bool {
	t.Helper()
	ok := errors.Is(err, target)
	switch {
	case ok:
		return true
	case err == nil:
		doError(t, name, args, "expected error %v, got nil", target)
	default:
		doError(t, name, args, "expected error %v, got %v", target, err)
	}
	return false
}

// AssertError fails the test if error is nil.
// Returns true if the assertion passed, false otherwise.
func AssertError(t T, err error, name string, args ...any) bool {
	t.Helper()
	ok := err != nil
	if !ok {
		doError(t, name, args, "expected an error but got nil")
	}
	return ok
}

// AssertNoError fails the test if error is not nil.
// Returns true if the assertion passed, false otherwise.
func AssertNoError(t T, err error, name string, args ...any) bool {
	t.Helper()
	ok := err == nil
	if !ok {
		doError(t, name, args, "unexpected error: %v", err)
	}
	return ok
}

// AssertContains fails the test if the string doesn't contain the substring.
// Returns true if the assertion passed, false otherwise.
func AssertContains(t T, str, substr string, name string, args ...any) bool {
	t.Helper()
	ok := strings.Contains(str, substr)
	if !ok {
		doError(t, name, args, "expected %q to contain %q", str, substr)
	}
	return ok
}

// AssertTypeIs fails the test if value is not of the expected type.
// It returns the value cast to the expected type and a boolean indicating success.
func AssertTypeIs[U any](t T, value any, name string, args ...any) (U, bool) {
	t.Helper()
	result, ok := value.(U)
	if !ok {
		var zero U
		doError(t, name, args, "expected type %T but got %T", zero, value)
		return zero, false
	}
	return result, true
}

// doError reports a test error with optional name prefix
func doError(t T, name string, nameArgs []any, msgFormat string, msgArgs ...any) {
	t.Helper()

	// Format the error message
	msg := fmt.Sprintf(msgFormat, msgArgs...)

	// Add name prefix if provided
	if name != "" {
		var prefix string
		if len(nameArgs) > 0 {
			prefix = fmt.Sprintf(name, nameArgs...)
		} else {
			prefix = name
		}
		msg = fmt.Sprintf("%s: %s", prefix, msg)
	}

	t.Errorf("%s", msg)
}
