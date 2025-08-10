# Common

[![Go Reference][godoc-badge]][godoc-link]
[![codecov][codecov-badge]][codecov-link]

Common provides shared utilities for protomcp.org projects.

## Overview

The `common` package serves as a foundation for other protomcp.org repositories,
providing reusable components for error handling and general utilities that are
shared between `nanorpc`, `protomcp`, and related projects.

## Features

- **Shared Types**: Reusable data structures and interfaces
- **Error Handling**: Consistent error patterns across projects
- **Utility Functions**: General-purpose helper functions like slice tools
- **Testing Support**: Integrates with `darvaza.org/core` for test utilities

## Slice Utilities

The common package provides generic slice manipulation functions designed to
prevent memory leaks when working with slices containing pointers or reference
types.

### ClearSlice

Zeros all elements in a slice (including unused capacity) and returns an empty
slice that reuses the same underlying array:

```go
// Example: Reusing a slice after clearing
responses := []Response{
    {ID: 1, Data: largeData1},
    {ID: 2, Data: largeData2},
    {ID: 3, Data: largeData3},
}

// Clear and reuse - all elements zeroed, length is 0, capacity preserved
responses = common.ClearSlice(responses)

// The slice can be reused without allocation
responses = append(responses, newResponse)
```

This is particularly useful for:

- Connection pools that reuse response slices
- Buffer management in high-throughput systems
- Preventing memory leaks from references in unused capacity

### ClearAndNilSlice

Zeros all elements in a slice (including unused capacity) and returns nil,
completely releasing the underlying array for garbage collection:

```go
// Example: Releasing memory completely
type Handler struct {
    buffer     []byte
    clients    []*Client
    responses  []Response
}

func (h *Handler) Close() {
    // Clear and release all memory
    h.buffer = common.ClearAndNilSlice(h.buffer)
    h.clients = common.ClearAndNilSlice(h.clients)
    h.responses = common.ClearAndNilSlice(h.responses)
}
```

Use cases:

- Clean-up in Close() or shutdown methods
- Releasing large temporary buffers
- Ensuring complete memory release for GC

### Memory Leak Prevention

These utilities solve a common Go pitfall where slicing doesn't clear elements
in the underlying array:

```go
// Problem: Memory leak with regular slicing
handlers := []*Handler{h1, h2, h3, h4, h5}
handlers = handlers[:2]  // h3, h4, h5 still referenced in underlying array!

// Solution: Use ClearSlice
handlers := []*Handler{h1, h2, h3, h4, h5}
handlers = common.ClearSlice(handlers)  // All pointers nil, no leaks
handlers = append(handlers, h1, h2)     // Safely reuse
```

## Usage

This package is designed to be imported by other protomcp.org projects:

```go
import (
    "protomcp.org/common"
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
