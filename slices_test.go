package common

import (
	"testing"

	"darvaza.org/core"
)

// Test types implementing TestCase interface

var _ core.TestCase = clearSliceTestCase{}
var _ core.TestCase = clearSlicePointerTestCase{}
var _ core.TestCase = clearAndNilSliceTestCase{}
var _ core.TestCase = clearAndNilSlicePointerTestCase{}

// clearSliceTestCase tests ClearSlice with int slices
type clearSliceTestCase struct {
	setup          func() []int
	name           string
	expectedCap    int
	checkNilResult bool // if true, expect non-nil result even from nil input
}

func (tc clearSliceTestCase) Name() string {
	return tc.name
}

func (tc clearSliceTestCase) Test(t *testing.T) {
	t.Helper()
	s := tc.setup()
	originalCap := cap(s)
	result := ClearSlice(s)

	// Check length - ClearSlice always returns zero length
	core.AssertEqual(t, 0, len(result), "length")

	// Check capacity
	expectedCap := tc.expectedCap
	if expectedCap == -1 {
		expectedCap = originalCap
	}
	core.AssertEqual(t, expectedCap, cap(result), "capacity")

	// Check nil/non-nil
	if tc.checkNilResult {
		core.AssertNotNil(t, result, "result")
	}

	// Verify all elements in capacity are zeroed
	if cap(result) > 0 {
		full := result[:cap(result)]
		for i, v := range full {
			core.AssertEqual(t, 0, v, "element[%d]", i)
		}
	}
}

func newClearSliceTestCase(name string, setup func() []int,
	expectedCap int, checkNilResult bool) clearSliceTestCase {
	return clearSliceTestCase{
		name:           name,
		setup:          setup,
		expectedCap:    expectedCap,
		checkNilResult: checkNilResult,
	}
}

// clearSlicePointerTestCase tests ClearSlice with pointer slices
type clearSlicePointerTestCase struct {
	setup       func() []*testNode
	name        string
	expectedLen int
	expectedCap int
}

func (tc clearSlicePointerTestCase) Name() string {
	return tc.name
}

func (tc clearSlicePointerTestCase) Test(t *testing.T) {
	t.Helper()
	s := tc.setup()
	originalCap := cap(s)
	result := ClearSlice(s)

	// Check length
	core.AssertEqual(t, tc.expectedLen, len(result), "length")

	// Check capacity
	expectedCap := tc.expectedCap
	if expectedCap == -1 {
		expectedCap = originalCap
	}
	core.AssertEqual(t, expectedCap, cap(result), "capacity")

	// Verify all pointers in capacity are nil
	if cap(result) > 0 {
		full := result[:cap(result)]
		for i, v := range full {
			core.AssertNil(t, v, "pointer[%d]", i)
		}
	}
}

func newClearSlicePointerTestCase(name string, setup func() []*testNode,
	expectedLen, expectedCap int) clearSlicePointerTestCase {
	return clearSlicePointerTestCase{
		name:        name,
		setup:       setup,
		expectedLen: expectedLen,
		expectedCap: expectedCap,
	}
}

// clearAndNilSliceTestCase tests ClearAndNilSlice with int slices
type clearAndNilSliceTestCase struct {
	setup              func() []int
	name               string
	checkOriginalArray bool // if true, verify the original array was cleared
}

func (tc clearAndNilSliceTestCase) Name() string {
	return tc.name
}

func (tc clearAndNilSliceTestCase) Test(t *testing.T) {
	t.Helper()
	s := tc.setup()

	// Keep reference to underlying array for validation
	var full []int
	if s != nil && tc.checkOriginalArray {
		full = s[:cap(s)]
	}

	result := ClearAndNilSlice(s)

	// Check result is nil - ClearAndNilSlice always returns nil
	core.AssertNil(t, result, "result")

	// Verify original underlying array was cleared
	if tc.checkOriginalArray && full != nil {
		for i, v := range full {
			core.AssertEqual(t, 0, v, "original[%d]", i)
		}
	}
}

func newClearAndNilSliceTestCase(name string, setup func() []int,
	checkOriginalArray bool) clearAndNilSliceTestCase {
	return clearAndNilSliceTestCase{
		name:               name,
		setup:              setup,
		checkOriginalArray: checkOriginalArray,
	}
}

// clearAndNilSlicePointerTestCase tests ClearAndNilSlice with pointer slices
type clearAndNilSlicePointerTestCase struct {
	setup              func() []*testNode
	name               string
	checkOriginalArray bool
}

func (tc clearAndNilSlicePointerTestCase) Name() string {
	return tc.name
}

func (tc clearAndNilSlicePointerTestCase) Test(t *testing.T) {
	t.Helper()
	s := tc.setup()

	// Keep reference to underlying array for validation
	var full []*testNode
	if s != nil && tc.checkOriginalArray {
		full = s[:cap(s)]
	}

	result := ClearAndNilSlice(s)

	// Check result is nil - ClearAndNilSlice always returns nil
	core.AssertNil(t, result, "result")

	// Verify all pointers in underlying array are nil
	if tc.checkOriginalArray && full != nil {
		for i, v := range full {
			core.AssertNil(t, v, "original[%d]", i)
		}
	}
}

func newClearAndNilSlicePointerTestCase(name string, setup func() []*testNode,
	checkOriginalArray bool) clearAndNilSlicePointerTestCase {
	return clearAndNilSlicePointerTestCase{
		name:               name,
		setup:              setup,
		checkOriginalArray: checkOriginalArray,
	}
}

// Test helpers

type testNode struct {
	Next  *testNode
	Value int
}

// Test case generators

func clearSliceBasicTestCases() []clearSliceTestCase {
	return []clearSliceTestCase{
		newClearSliceTestCase(
			"nil slice",
			func() []int { return nil },
			0,    // expectedCap
			true, // checkNilResult - expect non-nil result
		),
		newClearSliceTestCase(
			"empty slice",
			func() []int { return []int{} },
			0,     // expectedCap
			false, // checkNilResult
		),
		newClearSliceTestCase(
			"slice with values",
			func() []int { return []int{1, 2, 3} },
			3,     // expectedCap
			false, // checkNilResult
		),
	}
}

func clearSliceCapacityTestCases() []clearSliceTestCase {
	return []clearSliceTestCase{
		newClearSliceTestCase(
			"slice with spare capacity",
			func() []int {
				s := make([]int, 3, 10)
				s[0], s[1], s[2] = 1, 2, 3
				// Add more elements to underlying array
				extended := s[:6]
				extended[3], extended[4], extended[5] = 4, 5, 6
				return s
			},
			10,    // expectedCap
			false, // checkNilResult
		),
	}
}

func clearSlicePointerTestCases() []clearSlicePointerTestCase {
	return []clearSlicePointerTestCase{
		newClearSlicePointerTestCase(
			"slice with pointers",
			func() []*testNode {
				n1 := &testNode{Value: 1}
				n2 := &testNode{Value: 2}
				n3 := &testNode{Value: 3}
				return []*testNode{n1, n2, n3}
			},
			0,  // expectedLen
			-1, // expectedCap (-1 means use original capacity)
		),
	}
}

// Test functions

func TestClearSlice(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		core.RunTestCases(t, clearSliceBasicTestCases())
	})

	t.Run("capacity", func(t *testing.T) {
		core.RunTestCases(t, clearSliceCapacityTestCases())
	})

	t.Run("pointers", func(t *testing.T) {
		core.RunTestCases(t, clearSlicePointerTestCases())
	})

	t.Run("reuse after clear", func(t *testing.T) {
		s := []int{1, 2, 3}
		s = ClearSlice(s)

		// Should be able to append and reuse
		s = append(s, 10, 20)
		core.AssertEqual(t, 2, len(s), "length after append")
		core.AssertSliceEqual(t, []int{10, 20}, s, "values after append")
	})
}

// ClearAndNilSlice test cases

func clearAndNilSliceBasicTestCases() []clearAndNilSliceTestCase {
	return []clearAndNilSliceTestCase{
		newClearAndNilSliceTestCase(
			"nil slice",
			func() []int { return nil },
			false, // checkOriginalArray (no array to check)
		),
		newClearAndNilSliceTestCase(
			"empty slice",
			func() []int { return []int{} },
			false, // checkOriginalArray (empty)
		),
		newClearAndNilSliceTestCase(
			"slice with values",
			func() []int { return []int{1, 2, 3} },
			true, // checkOriginalArray
		),
	}
}

func clearAndNilSliceCapacityTestCases() []clearAndNilSliceTestCase {
	return []clearAndNilSliceTestCase{
		newClearAndNilSliceTestCase(
			"slice with spare capacity",
			func() []int {
				s := make([]int, 3, 10)
				s[0], s[1], s[2] = 1, 2, 3
				// Add more elements to underlying array
				extended := s[:6]
				extended[3], extended[4], extended[5] = 4, 5, 6
				return s
			},
			true, // checkOriginalArray
		),
	}
}

func clearAndNilSlicePointerTestCases() []clearAndNilSlicePointerTestCase {
	return []clearAndNilSlicePointerTestCase{
		newClearAndNilSlicePointerTestCase(
			"slice with pointers",
			func() []*testNode {
				n1 := &testNode{Value: 1}
				n2 := &testNode{Value: 2}
				n3 := &testNode{Value: 3}
				return []*testNode{n1, n2, n3}
			},
			true, // checkOriginalArray
		),
	}
}

func TestClearAndNilSlice(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		core.RunTestCases(t, clearAndNilSliceBasicTestCases())
	})

	t.Run("capacity", func(t *testing.T) {
		core.RunTestCases(t, clearAndNilSliceCapacityTestCases())
	})

	t.Run("pointers", func(t *testing.T) {
		core.RunTestCases(t, clearAndNilSlicePointerTestCases())
	})
}

// Benchmarks

type benchmarkSize struct {
	name string
	size int
}

func newBenchmarkSize(name string, size int) benchmarkSize {
	return benchmarkSize{
		name: name,
		size: size,
	}
}

func benchmarkSizes() []benchmarkSize {
	return []benchmarkSize{
		newBenchmarkSize("size-10", 10),
		newBenchmarkSize("size-100", 100),
		newBenchmarkSize("size-1000", 1000),
		newBenchmarkSize("size-10000", 10000),
	}
}

func runClearSliceBenchmark(b *testing.B, size int) {
	s := make([]int, size)
	for i := range s {
		s[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ClearSlice(s)
		s = s[:size] // restore length for next iteration
	}
}

func BenchmarkClearSlice(b *testing.B) {
	sizes := benchmarkSizes()
	for _, bs := range sizes {
		b.Run(bs.name, func(b *testing.B) {
			runClearSliceBenchmark(b, bs.size)
		})
	}
}

func runClearAndNilSliceBenchmark(b *testing.B, size int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, size)
		for j := range s {
			s[j] = j
		}
		_ = ClearAndNilSlice(s)
	}
}

func BenchmarkClearAndNilSlice(b *testing.B) {
	sizes := benchmarkSizes()
	for _, bs := range sizes {
		b.Run(bs.name, func(b *testing.B) {
			runClearAndNilSliceBenchmark(b, bs.size)
		})
	}
}
