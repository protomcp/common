package testutils

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type assertEqualTest struct {
	expected any
	actual   any
	name     string
	wantFail bool
}

func newAssertEqualTest(name string, expected, actual any, wantFail bool) assertEqualTest {
	return assertEqualTest{
		name:     name,
		expected: expected,
		actual:   actual,
		wantFail: wantFail,
	}
}

func (tt assertEqualTest) test(t *testing.T) {
	m := NewMockT()
	ok := AssertEqual(m, tt.expected, tt.actual, "")

	if tt.wantFail && ok {
		t.Error("expected assertion to fail but it passed")
	}
	if !tt.wantFail && !ok {
		t.Errorf("expected assertion to pass but it failed: %v", m.GetErrorf())
	}
}

func TestAssertEqual(t *testing.T) {
	tests := []assertEqualTest{
		newAssertEqualTest("equal ints", 42, 42, false),
		newAssertEqualTest("unequal ints", 42, 43, true),
		newAssertEqualTest("equal strings", "hello", "hello", false),
		newAssertEqualTest("unequal strings", "hello", "world", true),
		newAssertEqualTest("equal slices", []int{1, 2, 3}, []int{1, 2, 3}, false),
		newAssertEqualTest("unequal slices", []int{1, 2, 3}, []int{1, 2, 4}, true),
		newAssertEqualTest("nil slices", []int(nil), []int(nil), false),
		newAssertEqualTest("nil vs empty slice", []int(nil), []int{}, true),
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

func TestAssertEqualWithMessage(t *testing.T) {
	m := NewMockT()
	ok := AssertEqual(m, 1, 2, "values at index %d", 5)

	if ok {
		t.Fatal("expected assertion to fail")
	}

	errs := m.GetErrorf()
	if len(errs) != 1 {
		t.Fatalf("expected 1 error, got %d", len(errs))
	}

	if !strings.Contains(errs[0], "values at index 5") {
		t.Errorf("error message should contain custom message, got: %s", errs[0])
	}
}

type assertNotEqualTest struct {
	expected any
	actual   any
	name     string
	wantFail bool
}

func newAssertNotEqualTest(name string, expected, actual any, wantFail bool) assertNotEqualTest {
	return assertNotEqualTest{
		name:     name,
		expected: expected,
		actual:   actual,
		wantFail: wantFail,
	}
}

func (tt assertNotEqualTest) test(t *testing.T) {
	m := NewMockT()
	ok := AssertNotEqual(m, tt.expected, tt.actual, "")

	if tt.wantFail && ok {
		t.Error("expected assertion to fail but it passed")
	}
	if !tt.wantFail && !ok {
		t.Errorf("expected assertion to pass but it failed: %v", m.GetErrorf())
	}
}

func TestAssertNotEqual(t *testing.T) {
	tests := []assertNotEqualTest{
		newAssertNotEqualTest("equal ints", 42, 42, true),
		newAssertNotEqualTest("unequal ints", 42, 43, false),
		newAssertNotEqualTest("equal strings", "hello", "hello", true),
		newAssertNotEqualTest("unequal strings", "hello", "world", false),
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

type assertNilTest struct {
	value    any
	name     string
	wantFail bool
}

func newAssertNilTest(name string, value any, wantFail bool) assertNilTest {
	return assertNilTest{
		name:     name,
		value:    value,
		wantFail: wantFail,
	}
}

func (tt assertNilTest) test(t *testing.T) {
	m := NewMockT()
	ok := AssertNil(m, tt.value, "")

	if tt.wantFail && ok {
		t.Error("expected assertion to fail but it passed")
	}
	if !tt.wantFail && !ok {
		t.Errorf("expected assertion to pass but it failed: %v", m.GetErrorf())
	}
}

func TestAssertNil(t *testing.T) {
	tests := []assertNilTest{
		newAssertNilTest("nil pointer", (*int)(nil), false),
		newAssertNilTest("nil interface", any(nil), false),
		newAssertNilTest("nil slice", []int(nil), false),
		newAssertNilTest("nil map", map[string]int(nil), false),
		newAssertNilTest("non-nil pointer", new(int), true),
		newAssertNilTest("non-nil slice", []int{}, true),
		newAssertNilTest("non-nil map", make(map[string]int), true),
		newAssertNilTest("zero int", 0, true),
		newAssertNilTest("empty string", "", true),
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

type assertNotNilTest struct {
	value    any
	name     string
	wantFail bool
}

func newAssertNotNilTest(name string, value any, wantFail bool) assertNotNilTest {
	return assertNotNilTest{
		name:     name,
		value:    value,
		wantFail: wantFail,
	}
}

func (tt assertNotNilTest) test(t *testing.T) {
	m := NewMockT()
	ok := AssertNotNil(m, tt.value, "")

	if tt.wantFail && ok {
		t.Error("expected assertion to fail but it passed")
	}
	if !tt.wantFail && !ok {
		t.Errorf("expected assertion to pass but it failed: %v", m.GetErrorf())
	}
}

func TestAssertNotNil(t *testing.T) {
	tests := []assertNotNilTest{
		newAssertNotNilTest("nil pointer", (*int)(nil), true),
		newAssertNotNilTest("nil interface", any(nil), true),
		newAssertNotNilTest("nil slice", []int(nil), true),
		newAssertNotNilTest("non-nil pointer", new(int), false),
		newAssertNotNilTest("non-nil slice", []int{}, false),
		newAssertNotNilTest("zero int", 0, false),
		newAssertNotNilTest("empty string", "", false),
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

func TestAssertTrue(t *testing.T) {
	m := NewMockT()

	ok := AssertTrue(m, true, "")
	if !ok {
		t.Errorf("AssertTrue(true) should pass: %v", m.GetErrorf())
	}

	m.Reset()
	ok = AssertTrue(m, false, "")
	if ok {
		t.Error("AssertTrue(false) should fail")
	}
}

func TestAssertFalse(t *testing.T) {
	m := NewMockT()

	ok := AssertFalse(m, false, "")
	if !ok {
		t.Errorf("AssertFalse(false) should pass: %v", m.GetErrorf())
	}

	m.Reset()
	ok = AssertFalse(m, true, "")
	if ok {
		t.Error("AssertFalse(true) should fail")
	}
}

type assertErrorIsTest struct {
	err      error
	target   error
	name     string
	wantFail bool
}

func newAssertErrorIsTest(name string, err, target error, wantFail bool) assertErrorIsTest {
	return assertErrorIsTest{
		name:     name,
		err:      err,
		target:   target,
		wantFail: wantFail,
	}
}

func (tt assertErrorIsTest) test(t *testing.T) {
	t.Helper()
	m := NewMockT()
	ok := AssertErrorIs(m, tt.err, tt.target, "")

	if tt.wantFail && ok {
		t.Error("expected assertion to fail but it passed")
	}
	if !tt.wantFail && !ok {
		t.Errorf("expected assertion to pass but it failed: %v", m.GetErrorf())
	}
}

func makeAssertErrorIsTestCases() []assertErrorIsTest {
	errBase := errors.New("base error")
	errWrapped := fmt.Errorf("wrapped: %w", errBase)
	errOther := errors.New("other error")

	return []assertErrorIsTest{
		newAssertErrorIsTest("exact match", errBase, errBase, false),
		newAssertErrorIsTest("wrapped match", errWrapped, errBase, false),
		newAssertErrorIsTest("nil both", nil, nil, false),
		newAssertErrorIsTest("nil err non-nil target", nil, errBase, true),
		newAssertErrorIsTest("non-nil err nil target", errBase, nil, true),
		newAssertErrorIsTest("different errors", errBase, errOther, true),
		newAssertErrorIsTest("wrapped not matching", errWrapped, errOther, true),
	}
}

func TestAssertErrorIs(t *testing.T) {
	for _, tt := range makeAssertErrorIsTestCases() {
		t.Run(tt.name, tt.test)
	}
}

func TestAssertError(t *testing.T) {
	m := NewMockT()
	err := errors.New("test error")

	ok := AssertError(m, err, "")
	if !ok {
		t.Errorf("AssertError with error should pass: %v", m.GetErrorf())
	}

	m.Reset()
	ok = AssertError(m, nil, "")
	if ok {
		t.Error("AssertError with nil should fail")
	}
}

func TestAssertNoError(t *testing.T) {
	m := NewMockT()
	err := errors.New("test error")

	ok := AssertNoError(m, nil, "")
	if !ok {
		t.Errorf("AssertNoError with nil should pass: %v", m.GetErrorf())
	}

	m.Reset()
	ok = AssertNoError(m, err, "")
	if ok {
		t.Error("AssertNoError with error should fail")
	}
	errs := m.GetErrorf()
	if !strings.Contains(errs[0], "test error") {
		t.Errorf("error message should contain the error, got: %s", errs[0])
	}
}

type assertContainsTest struct {
	name     string
	str      string
	substr   string
	wantFail bool
}

func newAssertContainsTest(name, str, substr string, wantFail bool) assertContainsTest {
	return assertContainsTest{
		name:     name,
		str:      str,
		substr:   substr,
		wantFail: wantFail,
	}
}

func (tt assertContainsTest) test(t *testing.T) {
	m := NewMockT()
	ok := AssertContains(m, tt.str, tt.substr, "")

	if tt.wantFail && ok {
		t.Error("expected assertion to fail but it passed")
	}
	if !tt.wantFail && !ok {
		t.Errorf("expected assertion to pass but it failed: %v", m.GetErrorf())
	}
}

func TestAssertContains(t *testing.T) {
	tests := []assertContainsTest{
		newAssertContainsTest("contains substring", "hello world", "world", false),
		newAssertContainsTest("doesn't contain", "hello world", "foo", true),
		newAssertContainsTest("empty substring", "hello", "", false),
		newAssertContainsTest("both empty", "", "", false),
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

type assertTypeIsStringTest struct {
	value       any
	name        string
	expectValue string
	expectPass  bool
}

func (tc assertTypeIsStringTest) test(t *testing.T) {
	m := NewMockT()
	result, ok := AssertTypeIs[string](m, tc.value, "")

	if ok != tc.expectPass {
		t.Errorf("expected pass=%v, got %v", tc.expectPass, ok)
	}
	if result != tc.expectValue {
		t.Errorf("expected value %v, got %v", tc.expectValue, result)
	}
}

type assertTypeIsPointerTest struct {
	value       any
	expectValue *int
	name        string
	expectPass  bool
}

func (tc assertTypeIsPointerTest) test(t *testing.T) {
	m := NewMockT()
	result, ok := AssertTypeIs[*int](m, tc.value, "")

	if ok != tc.expectPass {
		t.Errorf("expected pass=%v, got %v", tc.expectPass, ok)
	}
	if result != tc.expectValue {
		t.Errorf("expected value %v, got %v", tc.expectValue, result)
	}
}

func testAssertTypeIsStrings(t *testing.T) {
	tests := []assertTypeIsStringTest{
		{
			name:        "correct type",
			value:       "hello",
			expectPass:  true,
			expectValue: "hello",
		},
		{
			name:        "incorrect type",
			value:       42,
			expectPass:  false,
			expectValue: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}
}

func testAssertTypeIsPointers(t *testing.T) {
	tests := []assertTypeIsPointerTest{
		{
			name:        "nil value",
			value:       nil,
			expectPass:  false,
			expectValue: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}
}

func TestAssertTypeIs(t *testing.T) {
	t.Run("string tests", testAssertTypeIsStrings)
	t.Run("pointer tests", testAssertTypeIsPointers)
}

func TestHelperCalls(t *testing.T) {
	m := NewMockT()

	_ = AssertEqual(m, 1, 2, "")
	// Both AssertEqual and doError call Helper()
	if m.GetHelperCalls() != 2 {
		t.Errorf("Helper() should be called twice, got %d calls", m.GetHelperCalls())
	}

	m.Reset()
	_ = AssertNil(m, 42, "")
	// Both AssertNil and doError call Helper()
	if m.GetHelperCalls() != 2 {
		t.Errorf("Helper() should be called twice, got %d calls", m.GetHelperCalls())
	}
}
