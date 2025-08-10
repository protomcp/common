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
package common
