package common

// ClearSlice zeros all elements in a slice (including unused capacity)
// and returns an empty slice that reuses the same underlying array.
// This prevents memory leaks when truncating slices containing pointers
// or other reference types.
//
// Example:
//
//	responses := []Response{{...}, {...}, {...}}
//	responses = ClearSlice(responses)  // All elements zeroed, length is 0
func ClearSlice[T any](s []T) []T {
	var zero T
	if s == nil {
		return []T{}
	}
	// Clear the entire capacity to ensure no data remains
	full := s[:cap(s)]
	for i := range full {
		full[i] = zero
	}
	return s[:0]
}

// ClearAndNilSlice zeros all elements in a slice (including unused capacity)
// and returns nil. This completely releases the underlying array for garbage
// collection.
//
// Example:
//
//	responses := []Response{{...}, {...}, {...}}
//	responses = ClearAndNilSlice(responses)  // All elements zeroed, slice is nil
func ClearAndNilSlice[T any](s []T) []T {
	if s != nil {
		var zero T
		// Clear the entire capacity to ensure no data remains
		full := s[:cap(s)]
		for i := range full {
			full[i] = zero
		}
	}
	return nil
}
