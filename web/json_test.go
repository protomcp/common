package web

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"darvaza.org/core"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Test types implementing TestCase interface

var _ core.TestCase = decodeJSONTestCase{}

// decodeJSONTestCase tests DecodeJSON function
type decodeJSONTestCase struct {
	createRequest   func() *http.Request
	out             func() *descriptorpb.DescriptorProto
	expectedName    string
	expectedPackage string
	name            string
	limit           int
	wantErr         bool
}

func (tc decodeJSONTestCase) Name() string {
	return tc.name
}

func (tc decodeJSONTestCase) Test(t *testing.T) {
	t.Helper()

	req := tc.createRequest()
	msg := tc.out()
	err := DecodeJSON(req, msg, tc.limit)

	if tc.wantErr {
		core.AssertError(t, err, "decode error")
		return
	}

	core.AssertNoError(t, err, "decode")

	if tc.expectedName != "" {
		core.AssertNotNil(t, msg.Name, "message name")
		core.AssertEqual(t, tc.expectedName, msg.GetName(), "name")
	}

	if tc.expectedPackage != "" {
		// For file descriptors, check package
		fileDesc := core.AssertMustTypeIs[*descriptorpb.FileDescriptorProto](t, msg, "file descriptor")
		core.AssertNotNil(t, fileDesc.Package, "package field")
		core.AssertEqual(t, tc.expectedPackage, fileDesc.GetPackage(), "package")
	}
}

//revive:disable-next-line:argument-limit
func newDecodeJSONTestCase(name string, createRequest func() *http.Request,
	out func() *descriptorpb.DescriptorProto, limit int, wantErr bool,
	expectedName, expectedPackage string) decodeJSONTestCase {
	return decodeJSONTestCase{
		name:            name,
		createRequest:   createRequest,
		out:             out,
		limit:           limit,
		wantErr:         wantErr,
		expectedName:    expectedName,
		expectedPackage: expectedPackage,
	}
}

func newDecodeJSONTestCaseSuccess(name string, createRequest func() *http.Request,
	out func() *descriptorpb.DescriptorProto, expectedName string) decodeJSONTestCase {
	return newDecodeJSONTestCase(name, createRequest, out, 0, false, expectedName, "")
}

func newDecodeJSONTestCaseError(name string, createRequest func() *http.Request,
	out func() *descriptorpb.DescriptorProto) decodeJSONTestCase {
	return newDecodeJSONTestCase(name, createRequest, out, 0, true, "", "")
}

func newDecodeJSONTestCaseLimit(name string, createRequest func() *http.Request,
	out func() *descriptorpb.DescriptorProto, limit int,
	wantErr bool) decodeJSONTestCase {
	return newDecodeJSONTestCase(name, createRequest, out, limit, wantErr, "", "")
}

// Test case generators

func decodeJSONBasicTestCases() []decodeJSONTestCase {
	validJSON := `{"name": "TestMessage"}`

	return []decodeJSONTestCase{
		newDecodeJSONTestCaseSuccess(
			"valid JSON",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(validJSON))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
			"TestMessage",
		),
		newDecodeJSONTestCaseError(
			"invalid JSON",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(`{"name": invalid}`))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
		),
		newDecodeJSONTestCaseSuccess(
			"empty JSON object",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader("{}"))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
			"",
		),
		newDecodeJSONTestCaseError(
			"empty request body",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(""))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
		),
	}
}

func decodeJSONLimitTestCases() []decodeJSONTestCase {
	smallJSON := `{"name": "Test"}`
	largeJSON := `{"name": "` + strings.Repeat("A", 1000) + `"}`

	return []decodeJSONTestCase{
		newDecodeJSONTestCaseLimit(
			"within default limit",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(smallJSON))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
			0, // use default
			false,
		),
		newDecodeJSONTestCaseLimit(
			"within custom limit",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(smallJSON))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
			100,
			false,
		),
		newDecodeJSONTestCaseLimit(
			"exceeds custom limit",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(largeJSON))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
			10,   // very small limit
			true, // should fail due to size limit
		),
		newDecodeJSONTestCaseLimit(
			"negative limit uses default",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(smallJSON))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
			-1,
			false,
		),
	}
}

func decodeJSONErrorTestCases() []decodeJSONTestCase {
	return []decodeJSONTestCase{
		newDecodeJSONTestCaseError(
			"read error from body",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", &errorReader{})
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
		),
		newDecodeJSONTestCaseError(
			"malformed JSON",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(`{`))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
		),
		newDecodeJSONTestCaseError(
			"wrong JSON type",
			func() *http.Request {
				req, _ := http.NewRequest("POST", "/test", strings.NewReader(`"string instead of object"`))
				return req
			},
			func() *descriptorpb.DescriptorProto { return &descriptorpb.DescriptorProto{} },
		),
	}
}

// Test helpers

type errorReader struct{}

func (*errorReader) Read([]byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

// Test functions

func TestDecodeJSON(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		core.RunTestCases(t, decodeJSONBasicTestCases())
	})

	t.Run("limits", func(t *testing.T) {
		core.RunTestCases(t, decodeJSONLimitTestCases())
	})

	t.Run("errors", func(t *testing.T) {
		core.RunTestCases(t, decodeJSONErrorTestCases())
	})

	t.Run("default limit constant", func(t *testing.T) {
		core.AssertEqual(t, 1024*1024, DefaultDecodeJSONLimit, "default limit")
	})

	t.Run("large payload within default limit", func(t *testing.T) {
		// Create a JSON payload just under the default limit
		payload := `{"name": "` + strings.Repeat("A", 1024*1024-20) + `"}`
		req, _ := http.NewRequest("POST", "/test", strings.NewReader(payload))
		msg := &descriptorpb.DescriptorProto{}

		err := DecodeJSON(req, msg, 0)
		core.AssertNoError(t, err, "large payload decode")
		core.AssertEqual(t, strings.Repeat("A", 1024*1024-20), msg.GetName(), "large name")
	})
}

// Benchmark tests

type decodeJSONBenchmarkCase struct {
	name    string
	payload string
	limit   int
}

func newDecodeJSONBenchmarkCase(name, payload string, limit int) decodeJSONBenchmarkCase {
	return decodeJSONBenchmarkCase{
		name:    name,
		payload: payload,
		limit:   limit,
	}
}

func decodeJSONBenchmarkCases() []decodeJSONBenchmarkCase {
	return []decodeJSONBenchmarkCase{
		newDecodeJSONBenchmarkCase("small", `{"name": "Test"}`, 0),
		newDecodeJSONBenchmarkCase("medium", `{"name": "`+strings.Repeat("A", 1000)+`"}`, 0),
		newDecodeJSONBenchmarkCase("large", `{"name": "`+strings.Repeat("A", 100000)+`"}`, 0),
	}
}

func BenchmarkDecodeJSON(b *testing.B) {
	cases := decodeJSONBenchmarkCases()
	for _, bc := range cases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				req, _ := http.NewRequest("POST", "/test", bytes.NewReader([]byte(bc.payload)))
				msg := &descriptorpb.DescriptorProto{}
				_ = DecodeJSON(req, msg, bc.limit)
			}
		})
	}
}
