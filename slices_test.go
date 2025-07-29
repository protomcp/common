package common

import (
	"testing"

	"protomcp.org/common/testutils"
)

// Test types for ClearSlice

type clearSliceTest struct {
	setup    func() []int
	validate func(t *testing.T, result []int)
	name     string
}

func (tc clearSliceTest) test(t *testing.T) {
	s := tc.setup()
	result := ClearSlice(s)
	tc.validate(t, result)
}

type clearSlicePointerTest struct {
	setup    func() []*testNode
	validate func(t *testing.T, result []*testNode)
	name     string
}

func (tc clearSlicePointerTest) test(t *testing.T) {
	s := tc.setup()
	result := ClearSlice(s)
	tc.validate(t, result)
}

// Test types for ClearAndNilSlice

type clearAndNilSliceTest struct {
	setup    func() []int
	validate func(t *testing.T, result, original []int)
	name     string
}

func (tc clearAndNilSliceTest) test(t *testing.T) {
	s := tc.setup()
	// Keep reference to underlying array for validation
	var full []int
	if s != nil {
		full = s[:cap(s)]
	}
	result := ClearAndNilSlice(s)
	tc.validate(t, result, full)
}

type clearAndNilSlicePointerTest struct {
	setup    func() []*testNode
	validate func(t *testing.T, result, original []*testNode)
	name     string
}

func (tc clearAndNilSlicePointerTest) test(t *testing.T) {
	s := tc.setup()
	// Keep reference to underlying array for validation
	var full []*testNode
	if s != nil {
		full = s[:cap(s)]
	}
	result := ClearAndNilSlice(s)
	tc.validate(t, result, full)
}

// Test helpers

type testNode struct {
	Next  *testNode
	Value int
}

func validateZeroedInts(t *testing.T, slice []int) {
	t.Helper()
	for i, v := range slice {
		testutils.AssertEqual(t, 0, v, "element[%d]", i)
	}
}

func validateNilPointers(t *testing.T, slice []*testNode) {
	t.Helper()
	for i, v := range slice {
		testutils.AssertNil(t, v, "element[%d]", i)
	}
}

// Test validation helpers

func validateEmptySlice(t *testing.T, result []int) {
	t.Helper()
	testutils.AssertEqual(t, 0, len(result), "length")
	testutils.AssertEqual(t, 0, cap(result), "capacity")
}

func validateClearedSlice(t *testing.T, result []int, expectedCap int) {
	t.Helper()
	testutils.AssertEqual(t, 0, len(result), "length")
	testutils.AssertEqual(t, expectedCap, cap(result), "capacity")
	// Verify all elements in capacity are zeroed
	full := result[:cap(result)]
	validateZeroedInts(t, full)
}

// Test case generators

func clearSliceBasicTests() []clearSliceTest {
	return []clearSliceTest{
		{
			name:  "nil slice",
			setup: func() []int { return nil },
			validate: func(t *testing.T, result []int) {
				testutils.AssertNotNil(t, result, "ClearSlice(nil) result")
				validateEmptySlice(t, result)
			},
		},
		{
			name:     "empty slice",
			setup:    func() []int { return []int{} },
			validate: validateEmptySlice,
		},
		{
			name:  "slice with values",
			setup: func() []int { return []int{1, 2, 3} },
			validate: func(t *testing.T, result []int) {
				validateClearedSlice(t, result, 3)
			},
		},
	}
}

// Named test functions

func testClearSliceBasic(t *testing.T) {
	for _, tc := range clearSliceBasicTests() {
		t.Run(tc.name, tc.test)
	}
}

func testClearSliceCapacity(t *testing.T) {
	tc := clearSliceTest{
		name: "slice with spare capacity",
		setup: func() []int {
			s := make([]int, 3, 10)
			s[0], s[1], s[2] = 1, 2, 3
			// Add more elements to underlying array
			extended := s[:6]
			extended[3], extended[4], extended[5] = 4, 5, 6
			return s
		},
		validate: func(t *testing.T, result []int) {
			testutils.AssertEqual(t, 0, len(result), "length")
			testutils.AssertEqual(t, 10, cap(result), "capacity")
			// Verify entire capacity is zeroed
			full := result[:cap(result)]
			validateZeroedInts(t, full)
		},
	}
	t.Run(tc.name, tc.test)
}

func testClearSlicePointers(t *testing.T) {
	tc := clearSlicePointerTest{
		name: "slice with pointers",
		setup: func() []*testNode {
			n1 := &testNode{Value: 1}
			n2 := &testNode{Value: 2}
			n3 := &testNode{Value: 3}
			return []*testNode{n1, n2, n3}
		},
		validate: func(t *testing.T, result []*testNode) {
			testutils.AssertEqual(t, 0, len(result), "length")
			// Verify all pointers are nil
			full := result[:cap(result)]
			validateNilPointers(t, full)
		},
	}
	t.Run(tc.name, tc.test)
}

func testClearSliceReuse(t *testing.T) {
	s := []int{1, 2, 3}
	s = ClearSlice(s)

	// Should be able to append and reuse
	s = append(s, 10, 20)
	testutils.AssertEqual(t, 2, len(s), "length after append")
	testutils.AssertEqual(t, []int{10, 20}, s, "values after append")
}

func TestClearSlice(t *testing.T) {
	t.Run("basic", testClearSliceBasic)
	t.Run("capacity", testClearSliceCapacity)
	t.Run("pointers", testClearSlicePointers)
	t.Run("reuse after clear", testClearSliceReuse)
}

// ClearAndNilSlice tests

func testClearAndNilSliceBasic(t *testing.T) {
	tests := []clearAndNilSliceTest{
		{
			name: "nil slice",
			setup: func() []int {
				return nil
			},
			validate: func(t *testing.T, result, _ []int) {
				testutils.AssertNil(t, result, "ClearAndNilSlice(nil) result")
			},
		},
		{
			name: "empty slice",
			setup: func() []int {
				return []int{}
			},
			validate: func(t *testing.T, result, _ []int) {
				testutils.AssertNil(t, result, "ClearAndNilSlice result")
			},
		},
		{
			name: "slice with values",
			setup: func() []int {
				return []int{1, 2, 3}
			},
			validate: func(t *testing.T, result, original []int) {
				testutils.AssertNil(t, result, "ClearAndNilSlice result")
				// Verify original underlying array was cleared
				validateZeroedInts(t, original)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}
}

func testClearAndNilSliceCapacity(t *testing.T) {
	tc := clearAndNilSliceTest{
		name: "slice with spare capacity",
		setup: func() []int {
			s := make([]int, 3, 10)
			s[0], s[1], s[2] = 1, 2, 3
			// Add more elements to underlying array
			extended := s[:6]
			extended[3], extended[4], extended[5] = 4, 5, 6
			return s
		},
		validate: func(t *testing.T, result []int, original []int) {
			testutils.AssertNil(t, result, "ClearAndNilSlice result")
			// Verify entire capacity was cleared
			validateZeroedInts(t, original)
		},
	}
	t.Run(tc.name, tc.test)
}

func testClearAndNilSlicePointers(t *testing.T) {
	tc := clearAndNilSlicePointerTest{
		name: "slice with pointers",
		setup: func() []*testNode {
			n1 := &testNode{Value: 1}
			n2 := &testNode{Value: 2}
			n3 := &testNode{Value: 3}
			return []*testNode{n1, n2, n3}
		},
		validate: func(t *testing.T, result, original []*testNode) {
			testutils.AssertNil(t, result, "ClearAndNilSlice result")
			// Verify all pointers in underlying array are nil
			validateNilPointers(t, original)
		},
	}
	t.Run(tc.name, tc.test)
}

func TestClearAndNilSlice(t *testing.T) {
	t.Run("basic", testClearAndNilSliceBasic)
	t.Run("capacity", testClearAndNilSliceCapacity)
	t.Run("pointers", testClearAndNilSlicePointers)
}

// Benchmarks

type benchmarkSize struct {
	name string
	size int
}

func newBenchmarkSizes() []benchmarkSize {
	return []benchmarkSize{
		{name: "size-10", size: 10},
		{name: "size-100", size: 100},
		{name: "size-1000", size: 1000},
		{name: "size-10000", size: 10000},
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
	sizes := newBenchmarkSizes()
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
	sizes := newBenchmarkSizes()
	for _, bs := range sizes {
		b.Run(bs.name, func(b *testing.B) {
			runClearAndNilSliceBenchmark(b, bs.size)
		})
	}
}
