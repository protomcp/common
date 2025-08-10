# Common

[![Go Reference][godoc-badge]][godoc-link]
[![Go Report Card][goreportcard-badge]][goreportcard-link]
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
- **Generator Utilities**: Common helpers for protoc plugin development in the
  `generator` submodule
- **Options System**: Flexible option override system for protoc plugins in the
  `options` submodule

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

## String Building with LazyBuffer

LazyBuffer is a convenience wrapper around strings.Builder that eliminates error
handling overhead. Unlike strings.Builder, LazyBuffer methods don't return
errors, making code cleaner and more readable.

### LazyBuffer Features

- **No error handling required**: Methods handle errors internally
- **Method chaining**: All write methods return the buffer for chaining
- **Nil-safe**: All methods safely handle nil receivers
- **Empty string filtering**: WriteString ignores empty strings
- **Full compatibility**: Built on strings.Builder for performance

### Basic Usage

```go
var buf common.LazyBuffer

// Simple string building
buf.WriteString("Hello", " ", "world")
fmt.Println(buf.String()) // "Hello world"

// Method chaining
output := buf.WriteString("func ").
    WriteString(name).
    WriteRunes('(').
    Printf("ctx %s", contextType).
    WriteRunes(')').
    String()
```

### Code Generation Example

LazyBuffer is particularly useful for code generation:

```go
func generateMethod(name string, params []Param) string {
    var buf common.LazyBuffer

    buf.WriteString("func (s *Service) ").
        WriteString(name).
        WriteRunes('(')

    for i, p := range params {
        if i > 0 {
            buf.WriteString(", ")
        }
        buf.Printf("%s %s", p.Name, p.Type)
    }

    return buf.WriteString(") error {\n").
        WriteString("\t// TODO: implement\n").
        WriteString("\treturn nil\n").
        WriteRunes('}').
        String()
}
```

### Template Rendering

Building complex strings without error handling clutter:

```go
func renderHTML(title string, items []Item) string {
    var buf common.LazyBuffer

    buf.Printf("<html>\n<head><title>%s</title></head>\n<body>\n", title).
        WriteString("<ul>\n")

    for _, item := range items {
        buf.Printf("  <li>%s: %s</li>\n", item.Name, item.Description)
    }

    return buf.WriteString("</ul>\n</body>\n</html>").String()
}
```

### Performance

LazyBuffer has minimal overhead compared to strings.Builder:

```go
// LazyBuffer - clean, chainable code
var buf common.LazyBuffer
result := buf.WriteString("Hello").
    WriteRunes(' ').
    WriteString("World").
    String()

// strings.Builder - requires error handling
var sb strings.Builder
_, _ = sb.WriteString("Hello")
_, _ = sb.WriteRune(' ')
_, _ = sb.WriteString("World")
result := sb.String()
```

The convenience comes with negligible performance cost, making it ideal for
cases where code clarity is more important than micro-optimizations.

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
[goreportcard-badge]: https://goreportcard.com/badge/protomcp.org/common
[goreportcard-link]: https://goreportcard.com/report/protomcp.org/common
[codecov-badge]: https://codecov.io/gh/protomcp/common/graph/badge.svg
[codecov-link]: https://codecov.io/gh/protomcp/common
