# Common

[![Go Reference][godoc-badge]][godoc-link]
[![codecov][codecov-badge]][codecov-link]

Common provides shared utilities and testing helpers for protomcp.org projects.

## Overview

The `common` package serves as a foundation for other protomcp.org repositories,
providing reusable components for testing, error handling, and general utilities
that are shared between `nanorpc`, `protomcp`, and related projects.

## Features

- **Test Utilities**: Common testing helpers and assertions
- **Shared Types**: Reusable data structures and interfaces
- **Error Handling**: Consistent error patterns across projects
- **Utility Functions**: General-purpose helper functions

## Test Utilities

The `testutils` package provides a foundation for testing across
protomcp.org projects:

### T Interface

A minimal interface that matches a subset of `testing.TB`:

```go
type T interface {
    Helper()
    Errorf(format string, args ...any)
    Fatalf(format string, args ...any)
}
```

This interface is satisfied by `*testing.T` and allows for creating test
helpers that work with both real tests and mock implementations.

### MockT

A test implementation of the T interface for testing test helpers:

```go
// Create a new MockT
m := testutils.NewMockT()

// Use it in place of *testing.T
MyTestHelper(m, someValue)

// Inspect what happened
errors := m.GetErrorf()        // []string of all Errorf calls
fatalMessages := m.GetFatalf() // []string of all Fatalf calls
helpers := m.GetHelperCalls()  // int count of Helper() calls

// Reset for reuse
m.Reset()
```

### Assertion Functions

The testutils package provides a comprehensive set of assertion functions that
all return bool, allowing for conditional logic:

```go
// Basic assertions
testutils.AssertEqual(t, expected, actual, "value")
testutils.AssertNotEqual(t, unexpected, actual, "value")

// Nil checking
testutils.AssertNil(t, ptr, "pointer should be nil")
testutils.AssertNotNil(t, result, "result required")

// Boolean assertions
testutils.AssertTrue(t, condition, "condition check")
testutils.AssertFalse(t, flag, "flag should be false")

// Error handling
testutils.AssertError(t, err, "operation should fail")
testutils.AssertNoError(t, err, "operation failed")
testutils.AssertErrorIs(t, err, ErrExpected, "wrong error type")

// String and type checking
testutils.AssertContains(t, output, "expected text", "output check")
value, ok := testutils.AssertTypeIs[*Config](t, result, "type assertion")
```

All assertions support formatted names:

```go
testutils.AssertEqual(t, 42, value, "item %d value", index)
```

### Helper Functions

```go
// IsNil checks if a value is nil using reflection
if testutils.IsNil(value) {
    // handle nil case
}
```

## Usage

This package is designed to be imported by other protomcp.org projects:

```go
import (
    "protomcp.org/common"
    "protomcp.org/common/testutils"
)
```

## Development

For development guidelines, please refer to [AGENT.md](AGENT.md).

## License

See [LICENCE.txt](LICENCE.txt) for licensing information.

[godoc-badge]: https://pkg.go.dev/badge/protomcp.org/common.svg
[godoc-link]: https://pkg.go.dev/protomcp.org/common
[codecov-badge]: https://codecov.io/gh/protomcp/common/graph/badge.svg
[codecov-link]: https://codecov.io/gh/protomcp/common
