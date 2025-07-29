package testutils_test

import (
	"testing"

	"protomcp.org/common/testutils"
)

type nilTestCase struct {
	value any
	name  string
	want  bool
}

func newNilTestCase(name string, value any, want bool) nilTestCase {
	return nilTestCase{
		value: value,
		name:  name,
		want:  want,
	}
}

func (tc nilTestCase) test(t *testing.T) {
	t.Helper()
	got := testutils.IsNil(tc.value)
	if got != tc.want {
		t.Errorf("IsNil(%v) = %v, want %v", tc.value, got, tc.want)
	}
}

func makeIsNilTestCases() []nilTestCase {
	var nilPtr *int
	var nilSlice []int
	var nilMap map[string]int
	var nilChan chan int
	var nilFunc func()
	var nilInterface any

	nonNilPtr := new(int)
	nonNilSlice := []int{}
	nonNilMap := make(map[string]int)
	nonNilChan := make(chan int)
	nonNilFunc := func() {}

	return []nilTestCase{
		// Nil values
		newNilTestCase("nil literal", nil, true),
		newNilTestCase("nil pointer", nilPtr, true),
		newNilTestCase("nil slice", nilSlice, true),
		newNilTestCase("nil map", nilMap, true),
		newNilTestCase("nil channel", nilChan, true),
		newNilTestCase("nil function", nilFunc, true),
		newNilTestCase("nil interface", nilInterface, true),

		// Non-nil values
		newNilTestCase("non-nil pointer", nonNilPtr, false),
		newNilTestCase("non-nil slice", nonNilSlice, false),
		newNilTestCase("non-nil map", nonNilMap, false),
		newNilTestCase("non-nil channel", nonNilChan, false),
		newNilTestCase("non-nil function", nonNilFunc, false),
		newNilTestCase("string", "hello", false),
		newNilTestCase("number", 42, false),
		newNilTestCase("bool", true, false),
		newNilTestCase("struct", struct{}{}, false),

		// Special cases
		newNilTestCase("typed nil interface", (*int)(nil), true),
		newNilTestCase("empty string", "", false),
		newNilTestCase("zero int", 0, false),
		newNilTestCase("false bool", false, false),
	}
}

func TestIsNil(t *testing.T) {
	for _, tt := range makeIsNilTestCases() {
		t.Run(tt.name, tt.test)
	}
}
