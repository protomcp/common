// Package common provides shared utilities and helpers for protomcp.org projects.
//
// The common package serves as a foundation library for the protomcp.org ecosystem,
// preventing code duplication across repositories like nanorpc and protomcp. It
// provides general-purpose utilities, test helpers, and patterns that are shared
// across multiple projects.
//
// # Slice Utilities
//
// The package provides generic slice manipulation functions that help prevent
// memory leaks when working with slices containing pointers or reference types:
//
//   - ClearSlice: Zeros all elements including unused capacity and returns an
//     empty slice that reuses the underlying array
//   - ClearAndNilSlice: Zeros all elements and returns nil, releasing the
//     underlying array for garbage collection
//
// These utilities are particularly useful when truncating slices that contain
// pointers, interfaces, or other reference types, as they ensure no references
// remain in the unused capacity that could prevent garbage collection.
//
// Example usage:
//
//	// Prevent memory leaks when clearing a slice of pointers
//	handlers := []*Handler{h1, h2, h3}
//	handlers = ClearSlice(handlers)  // All pointers nil, slice empty but reusable
//
//	// Or completely release the underlying array
//	responses := []Response{{...}, {...}, {...}}
//	responses = ClearAndNilSlice(responses)  // All cleared, slice is nil
//
// # String Building
//
// LazyBuffer provides a convenient wrapper around strings.Builder that eliminates
// the need for error handling. It's particularly useful for code generation,
// template rendering, or any scenario where you're building strings incrementally:
//
//   - WriteString: Appends strings, ignoring empty ones
//   - WriteRunes: Appends individual runes
//   - Printf: Appends formatted strings
//   - All methods support chaining and nil-safety
//
// Example usage:
//
//	var buf LazyBuffer
//	output := buf.WriteString("func ").
//		WriteString(name).
//		WriteRunes('(').
//		Printf("ctx %s", contextType).
//		WriteRunes(')').
//		String()
package common
