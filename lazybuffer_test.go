package common

import (
	"strings"
	"testing"

	"darvaza.org/core"
)

// Test types for LazyBuffer

type lazyBufferTest struct {
	name     string
	setup    func() *LazyBuffer
	action   func(*LazyBuffer)
	expected string
}

func (tc lazyBufferTest) test(t *testing.T) {
	buf := tc.setup()
	tc.action(buf)
	core.AssertEqual(t, tc.expected, buf.String(), tc.name)
}

type lazyBufferChainTest struct {
	name     string
	chain    func(*LazyBuffer) string
	expected string
}

func (tc lazyBufferChainTest) test(t *testing.T) {
	var buf LazyBuffer
	result := tc.chain(&buf)
	core.AssertEqual(t, tc.expected, result, tc.name)
}

type lazyBufferNilTest struct {
	action func(*LazyBuffer) any
	check  func(t *testing.T, result any)
	name   string
}

func newLazyBufferNilTest(
	name string,
	action func(*LazyBuffer) any,
	check func(t *testing.T, result any),
) lazyBufferNilTest {
	return lazyBufferNilTest{
		name:   name,
		action: action,
		check:  check,
	}
}

func (tc lazyBufferNilTest) test(t *testing.T) {
	var buf *LazyBuffer
	result := tc.action(buf)
	tc.check(t, result)
}

// Test functions

func testLazyBufferWriteString(t *testing.T) {
	tests := []lazyBufferTest{
		{
			name:  "single string",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteString("hello")
			},
			expected: "hello",
		},
		{
			name:  "multiple strings",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteString("hello", " ", "world")
			},
			expected: "hello world",
		},
		{
			name:  "empty strings ignored",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteString("", "hello", "", "", "world", "")
			},
			expected: "helloworld",
		},
		{
			name:  "all empty strings",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteString("", "", "")
			},
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}
}

func testLazyBufferWriteRunes(t *testing.T) {
	tests := []lazyBufferTest{
		{
			name:  "ASCII runes",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteRunes('H', 'e', 'l', 'l', 'o')
			},
			expected: "Hello",
		},
		{
			name:  "Unicode runes",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteRunes('世', '界', '🌍')
			},
			expected: "世界🌍",
		},
		{
			name:  "mixed ASCII and Unicode",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteRunes('H', 'i', ' ', '世', '界')
			},
			expected: "Hi 世界",
		},
		{
			name:  "single rune",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteRunes('!')
			},
			expected: "!",
		},
		{
			name:  "no runes",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.WriteRunes()
			},
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}
}

func testLazyBufferPrintf(t *testing.T) {
	tests := []lazyBufferTest{
		{
			name:  "simple format",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.Printf("Hello %s", "world")
			},
			expected: "Hello world",
		},
		{
			name:  "multiple format verbs",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.Printf("Name: %s, Age: %d, Score: %.2f", "Alice", 30, 95.5)
			},
			expected: "Name: Alice, Age: 30, Score: 95.50",
		},
		{
			name:  "no format verbs",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.Printf("Just a plain string")
			},
			expected: "Just a plain string",
		},
		{
			name:  "empty format",
			setup: func() *LazyBuffer { return new(LazyBuffer) },
			action: func(b *LazyBuffer) {
				b.Printf("")
			},
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}
}

func testLazyBufferChaining(t *testing.T) {
	tests := []lazyBufferChainTest{
		{
			name: "full chain",
			chain: func(b *LazyBuffer) string {
				return b.WriteString("Hello").
					WriteRunes(' ').
					WriteString("world").
					Printf(", %d", 42).
					WriteRunes('!').
					String()
			},
			expected: "Hello world, 42!",
		},
		{
			name: "chain with empty strings",
			chain: func(b *LazyBuffer) string {
				return b.WriteString("", "Start").
					WriteString("", "").
					WriteRunes('-').
					WriteString("End", "").
					String()
			},
			expected: "Start-End",
		},
		{
			name: "complex chain",
			chain: func(b *LazyBuffer) string {
				return b.Printf("Line %d: ", 1).
					WriteString("Error").
					WriteRunes(':', ' ').
					Printf("%q", "file not found").
					WriteRunes('\n').
					String()
			},
			expected: "Line 1: Error: \"file not found\"\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}
}

func testLazyBufferNil(t *testing.T) {
	tests := []lazyBufferNilTest{
		newLazyBufferNilTest(
			"nil WriteString returns nil",
			func(b *LazyBuffer) any { return b.WriteString("test") },
			func(t *testing.T, result any) {
				core.AssertNil(t, result, "WriteString on nil buffer")
			},
		),
		newLazyBufferNilTest(
			"nil WriteRunes returns nil",
			func(b *LazyBuffer) any { return b.WriteRunes('a', 'b') },
			func(t *testing.T, result any) {
				core.AssertNil(t, result, "WriteRunes on nil buffer")
			},
		),
		newLazyBufferNilTest(
			"nil Printf returns nil",
			func(b *LazyBuffer) any { return b.Printf("test %d", 1) },
			func(t *testing.T, result any) {
				core.AssertNil(t, result, "Printf on nil buffer")
			},
		),
		newLazyBufferNilTest(
			"nil String returns empty",
			func(b *LazyBuffer) any { return b.String() },
			func(t *testing.T, result any) {
				core.AssertEqual(t, "", result, "String on nil buffer")
			},
		),
		newLazyBufferNilTest(
			"nil Len returns 0",
			func(b *LazyBuffer) any { return b.Len() },
			func(t *testing.T, result any) {
				core.AssertEqual(t, 0, result, "Len on nil buffer")
			},
		),
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.test)
	}

	// Test Reset on nil doesn't panic
	t.Run("nil Reset no panic", func(t *testing.T) {
		var buf *LazyBuffer
		err := core.Catch(func() error {
			buf.Reset()
			return nil
		})
		core.AssertNoError(t, err, "Reset on nil buffer")
	})
}

func testLazyBufferLen(t *testing.T) {
	var buf LazyBuffer

	core.AssertEqual(t, 0, buf.Len(), "initial length")

	buf.WriteString("Hello")
	core.AssertEqual(t, 5, buf.Len(), "length after Hello")

	buf.WriteRunes(' ', '世', '界')
	// "Hello 世界" - 5 + 1 + 3 + 3 = 12 bytes (世 and 界 are 3 bytes each in UTF-8)
	core.AssertEqual(t, 12, buf.Len(), "length after Unicode")

	buf.Printf(" %d", 42)
	core.AssertEqual(t, 15, buf.Len(), "length after Printf")
}

func testLazyBufferReset(t *testing.T) {
	var buf LazyBuffer

	buf.WriteString("Hello world")
	core.AssertEqual(t, "Hello world", buf.String(), "before reset")
	core.AssertEqual(t, 11, buf.Len(), "length before reset")

	buf.Reset()
	core.AssertEqual(t, "", buf.String(), "after reset")
	core.AssertEqual(t, 0, buf.Len(), "length after reset")

	// Should be able to use after reset
	buf.WriteString("New content")
	core.AssertEqual(t, "New content", buf.String(), "after reset and write")
}

func testLazyBufferLargeContent(t *testing.T) {
	var buf LazyBuffer

	// Build a large string
	const iterations = 1000
	for i := range iterations {
		buf.Printf("Line %d: This is some test content\n", i)
	}

	result := buf.String()
	lines := strings.Count(result, "\n")
	core.AssertEqual(t, iterations, lines, "number of lines")

	// Verify it starts and ends correctly
	core.AssertTrue(t, strings.HasPrefix(result, "Line 0:"), "starts with Line 0")
	core.AssertTrue(t, strings.Contains(result, "Line 999:"), "contains Line 999")
}

func TestLazyBuffer(t *testing.T) {
	t.Run("WriteString", testLazyBufferWriteString)
	t.Run("WriteRunes", testLazyBufferWriteRunes)
	t.Run("Printf", testLazyBufferPrintf)
	t.Run("chaining", testLazyBufferChaining)
	t.Run("nil safety", testLazyBufferNil)
	t.Run("Len", testLazyBufferLen)
	t.Run("Reset", testLazyBufferReset)
	t.Run("large content", testLazyBufferLargeContent)
}

// Benchmarks

func benchmarkWriteString(b *testing.B) {
	for range b.N {
		var buf LazyBuffer
		buf.WriteString("Hello", " ", "world", " ", "from", " ", "LazyBuffer")
		_ = buf.String()
	}
}

func benchmarkWriteRunes(b *testing.B) {
	for range b.N {
		var buf LazyBuffer
		buf.WriteRunes('H', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd')
		_ = buf.String()
	}
}

func benchmarkPrintf(b *testing.B) {
	for i := range b.N {
		var buf LazyBuffer
		buf.Printf("Item %d: %s (%.2f%%)", i, "test", 99.9)
		_ = buf.String()
	}
}

func benchmarkChained(b *testing.B) {
	for i := range b.N {
		var buf LazyBuffer
		_ = buf.WriteString("Start").
			WriteRunes('-').
			Printf("%d", i).
			WriteRunes('-').
			WriteString("End").
			String()
	}
}

func BenchmarkLazyBuffer(b *testing.B) {
	b.Run("WriteString", benchmarkWriteString)
	b.Run("WriteRunes", benchmarkWriteRunes)
	b.Run("Printf", benchmarkPrintf)
	b.Run("chained", benchmarkChained)
}

func benchmarkLazyBufferLines(b *testing.B) {
	for range b.N {
		var buf LazyBuffer
		for j := range 100 {
			buf.Printf("Line %d\n", j)
		}
		_ = buf.String()
	}
}

func benchmarkStringsBuilderLines(b *testing.B) {
	for range b.N {
		var buf strings.Builder
		for j := range 100 {
			_, _ = buf.WriteString("Line ")
			_, _ = buf.WriteString(string(rune('0' + j/10)))
			_, _ = buf.WriteString(string(rune('0' + j%10)))
			_, _ = buf.WriteRune('\n')
		}
		_ = buf.String()
	}
}

func BenchmarkLazyBufferVsStringsBuilder(b *testing.B) {
	b.Run("LazyBuffer", benchmarkLazyBufferLines)
	b.Run("strings.Builder", benchmarkStringsBuilderLines)
}
