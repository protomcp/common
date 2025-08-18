package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"darvaza.org/core"
	"darvaza.org/x/web/consts"
	"darvaza.org/x/web/resource"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Test types implementing TestCase interface

var _ core.TestCase = renderProtoJSONTestCase{}
var _ core.TestCase = renderJSONTestCase{}
var _ core.TestCase = doRenderJSONTestCase{}

// renderProtoJSONTestCase tests RenderProtoJSON function
type renderProtoJSONTestCase struct {
	createRequest func() (http.ResponseWriter, *http.Request)
	data          proto.Message
	expectedJSON  string
	name          string
	wantErr       bool
	checkHeaders  bool
}

func (tc renderProtoJSONTestCase) Name() string {
	return tc.name
}

func (tc renderProtoJSONTestCase) Test(t *testing.T) {
	t.Helper()

	rw, req := tc.createRequest()
	err := RenderProtoJSON(rw, req, tc.data)

	// Convert back to *httptest.ResponseRecorder for testing
	var recorder *httptest.ResponseRecorder
	if rw != nil {
		recorder = core.AssertMustTypeIs[*httptest.ResponseRecorder](t, rw, "response writer")
	}

	tc.testResult(t, recorder, req, err)
}

func (tc renderProtoJSONTestCase) testResult(t *testing.T, rw *httptest.ResponseRecorder,
	req *http.Request, err error) {
	t.Helper()

	if tc.wantErr {
		core.AssertError(t, err, "render error")
		return
	}

	core.AssertNoError(t, err, "render")

	if tc.checkHeaders {
		tc.testHeaders(t, rw, req)
	}

	tc.testBody(t, rw, req)
}

func (tc renderProtoJSONTestCase) testHeaders(t *testing.T, rw *httptest.ResponseRecorder, req *http.Request) {
	t.Helper()

	if rw == nil {
		return
	}

	core.AssertEqual(t, consts.JSON, rw.Header().Get(consts.ContentType), "content-type")

	if req != nil && req.Method != consts.HEAD {
		contentLength := rw.Header().Get(consts.ContentLength)
		core.AssertNotEqual(t, "", contentLength, "content-length header")
	}
}

func (tc renderProtoJSONTestCase) testBody(t *testing.T, rw *httptest.ResponseRecorder, req *http.Request) {
	t.Helper()

	if rw == nil || req == nil {
		return
	}

	if req.Method == consts.HEAD {
		core.AssertEqual(t, "", rw.Body.String(), "HEAD response body")
		return
	}

	if tc.expectedJSON != "" {
		core.AssertEqual(t, tc.expectedJSON, rw.Body.String(), "response body")
	}
}

//revive:disable-next-line:argument-limit
func newRenderProtoJSONTestCase(name string,
	createRequest func() (http.ResponseWriter, *http.Request),
	data proto.Message, wantErr bool, expectedJSON string,
	checkHeaders bool) renderProtoJSONTestCase {
	return renderProtoJSONTestCase{
		name:          name,
		createRequest: createRequest,
		data:          data,
		wantErr:       wantErr,
		expectedJSON:  expectedJSON,
		checkHeaders:  checkHeaders,
	}
}

func newRenderProtoJSONTestCaseSuccess(name string,
	createRequest func() (http.ResponseWriter, *http.Request),
	data proto.Message, expectedJSON string) renderProtoJSONTestCase {
	return newRenderProtoJSONTestCase(name, createRequest, data, false, expectedJSON, true)
}

func newRenderProtoJSONTestCaseError(name string,
	createRequest func() (http.ResponseWriter, *http.Request),
	data proto.Message) renderProtoJSONTestCase {
	return newRenderProtoJSONTestCase(name, createRequest, data, true, "", false)
}

// renderJSONTestCase tests RenderJSON function
type renderJSONTestCase struct {
	createRequest func() (http.ResponseWriter, *http.Request)
	data          any
	expectedJSON  string
	name          string
	wantErr       bool
	checkHeaders  bool
}

func (tc renderJSONTestCase) Name() string {
	return tc.name
}

func (tc renderJSONTestCase) Test(t *testing.T) {
	t.Helper()

	rw, req := tc.createRequest()
	err := RenderJSON(rw, req, tc.data)

	// Convert back to *httptest.ResponseRecorder for testing
	var recorder *httptest.ResponseRecorder
	if rw != nil {
		recorder = core.AssertMustTypeIs[*httptest.ResponseRecorder](t, rw, "response writer")
	}

	tc.testResult(t, recorder, req, err)
}

func (tc renderJSONTestCase) testResult(t *testing.T, rw *httptest.ResponseRecorder, req *http.Request, err error) {
	t.Helper()

	if tc.wantErr {
		core.AssertError(t, err, "render error")
		return
	}

	core.AssertNoError(t, err, "render")

	if tc.checkHeaders {
		tc.testHeaders(t, rw, req)
	}

	tc.testBody(t, rw, req)
}

func (tc renderJSONTestCase) testHeaders(t *testing.T, rw *httptest.ResponseRecorder, req *http.Request) {
	t.Helper()

	if rw == nil {
		return
	}

	core.AssertEqual(t, consts.JSON, rw.Header().Get(consts.ContentType), "content-type")

	if req != nil && req.Method != consts.HEAD {
		contentLength := rw.Header().Get(consts.ContentLength)
		core.AssertNotEqual(t, "", contentLength, "content-length header")
	}
}

func (tc renderJSONTestCase) testBody(t *testing.T, rw *httptest.ResponseRecorder, req *http.Request) {
	t.Helper()

	if rw == nil || req == nil {
		return
	}

	if req.Method == consts.HEAD {
		core.AssertEqual(t, "", rw.Body.String(), "HEAD response body")
		return
	}

	if tc.expectedJSON != "" {
		core.AssertEqual(t, tc.expectedJSON, rw.Body.String(), "response body")
	}
}

//revive:disable-next-line:argument-limit
func newRenderJSONTestCase(name string,
	createRequest func() (http.ResponseWriter, *http.Request),
	data any, wantErr bool, expectedJSON string,
	checkHeaders bool) renderJSONTestCase {
	return renderJSONTestCase{
		name:          name,
		createRequest: createRequest,
		data:          data,
		wantErr:       wantErr,
		expectedJSON:  expectedJSON,
		checkHeaders:  checkHeaders,
	}
}

func newRenderJSONTestCaseSuccess(name string,
	createRequest func() (http.ResponseWriter, *http.Request),
	data any, expectedJSON string) renderJSONTestCase {
	return newRenderJSONTestCase(name, createRequest, data, false, expectedJSON, true)
}

func newRenderJSONTestCaseError(name string,
	createRequest func() (http.ResponseWriter, *http.Request),
	data any) renderJSONTestCase {
	return newRenderJSONTestCase(name, createRequest, data, true, "", false)
}

// doRenderJSONTestCase tests the internal doRenderJSON function
type doRenderJSONTestCase struct {
	createRequest func() (http.ResponseWriter, *http.Request)
	marshalFunc   func(any) ([]byte, error)
	data          any
	name          string
	wantErr       bool
}

func (tc doRenderJSONTestCase) Name() string {
	return tc.name
}

func (tc doRenderJSONTestCase) Test(t *testing.T) {
	t.Helper()

	rw, req := tc.createRequest()
	err := doRenderJSON(rw, req, tc.marshalFunc, tc.data)

	if tc.wantErr {
		core.AssertError(t, err, "render error")
		return
	}

	core.AssertNoError(t, err, "render")

	// Convert back to *httptest.ResponseRecorder for testing
	core.AssertMustNotNil(t, rw, "response writer")
	recorder := core.AssertMustTypeIs[*httptest.ResponseRecorder](t, rw, "response writer")
	core.AssertEqual(t, consts.JSON, recorder.Header().Get(consts.ContentType), "content-type")
}

func newDoRenderJSONTestCase(name string,
	createRequest func() (http.ResponseWriter, *http.Request),
	marshalFunc func(any) ([]byte, error), data any,
	wantErr bool) doRenderJSONTestCase {
	return doRenderJSONTestCase{
		name:          name,
		createRequest: createRequest,
		marshalFunc:   marshalFunc,
		data:          data,
		wantErr:       wantErr,
	}
}

// Test case generators

func renderProtoJSONBasicTestCases() []renderProtoJSONTestCase {
	testMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("TestMessage"),
	}

	return []renderProtoJSONTestCase{
		newRenderProtoJSONTestCaseSuccess(
			"valid protobuf message",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("POST", "/test", nil)
			},
			testMsg,
			`{"name":"TestMessage"}`,
		),
		newRenderProtoJSONTestCaseSuccess(
			"nil message",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("POST", "/test", nil)
			},
			nil,
			"{}",
		),
		newRenderProtoJSONTestCaseSuccess(
			"HEAD request",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("HEAD", "/test", nil)
			},
			testMsg,
			"", // no body expected for HEAD
		),
	}
}

func renderProtoJSONErrorTestCases() []renderProtoJSONTestCase {
	return []renderProtoJSONTestCase{
		newRenderProtoJSONTestCaseError(
			"nil response writer",
			func() (http.ResponseWriter, *http.Request) {
				return nil, httptest.NewRequest("POST", "/test", nil)
			},
			&descriptorpb.DescriptorProto{},
		),
		newRenderProtoJSONTestCaseError(
			"nil request",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), nil
			},
			&descriptorpb.DescriptorProto{},
		),
	}
}

func renderJSONBasicTestCases() []renderJSONTestCase {
	testData := map[string]string{
		"key": "value",
	}

	return []renderJSONTestCase{
		newRenderJSONTestCaseSuccess(
			"valid data",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("POST", "/test", nil)
			},
			testData,
			`{"key":"value"}`,
		),
		newRenderJSONTestCaseSuccess(
			"nil data",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("POST", "/test", nil)
			},
			nil,
			"null",
		),
		newRenderJSONTestCaseSuccess(
			"HEAD request",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("HEAD", "/test", nil)
			},
			testData,
			"", // no body expected for HEAD
		),
	}
}

func renderJSONErrorTestCases() []renderJSONTestCase {
	return []renderJSONTestCase{
		newRenderJSONTestCaseError(
			"nil response writer",
			func() (http.ResponseWriter, *http.Request) {
				return nil, httptest.NewRequest("POST", "/test", nil)
			},
			map[string]string{},
		),
		newRenderJSONTestCaseError(
			"nil request",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), nil
			},
			map[string]string{},
		),
		newRenderJSONTestCaseError(
			"unmarshallable data",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("POST", "/test", nil)
			},
			make(chan int), // channels cannot be marshalled to JSON
		),
	}
}

func doRenderJSONTestCases() []doRenderJSONTestCase {
	return []doRenderJSONTestCase{
		newDoRenderJSONTestCase(
			"nil marshal function",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("POST", "/test", nil)
			},
			nil,
			"test",
			true,
		),
		newDoRenderJSONTestCase(
			"marshal error",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("POST", "/test", nil)
			},
			func(any) ([]byte, error) {
				return nil, &json.UnsupportedTypeError{}
			},
			"test",
			true,
		),
		newDoRenderJSONTestCase(
			"successful marshal",
			func() (http.ResponseWriter, *http.Request) {
				return httptest.NewRecorder(), httptest.NewRequest("POST", "/test", nil)
			},
			json.Marshal,
			"test",
			false,
		),
	}
}

// Test functions

func TestRenderProtoJSON(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		core.RunTestCases(t, renderProtoJSONBasicTestCases())
	})

	t.Run("errors", func(t *testing.T) {
		core.RunTestCases(t, renderProtoJSONErrorTestCases())
	})

	t.Run("enum serialization", func(t *testing.T) {
		// Test that enums are serialized as strings not integers
		field := &descriptorpb.FieldDescriptorProto{
			Name: proto.String("test_field"),
			Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
		}

		rw := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", nil)

		err := RenderProtoJSON(rw, req, field)
		core.AssertNoError(t, err, "render")

		body := rw.Body.String()
		core.AssertContains(t, body, `"TYPE_STRING"`, "enum as string")
		// Verify it doesn't contain the integer representation
		core.AssertFalse(t, strings.Contains(body, `"type":9`), "enum not as integer")
	})
}

func TestRenderJSON(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		core.RunTestCases(t, renderJSONBasicTestCases())
	})

	t.Run("errors", func(t *testing.T) {
		core.RunTestCases(t, renderJSONErrorTestCases())
	})
}

func TestDoRenderJSON(t *testing.T) {
	core.RunTestCases(t, doRenderJSONTestCases())
}

func TestWithProtoJSON(t *testing.T) {
	t.Run("option creation", func(t *testing.T) {
		option := WithProtoJSON[*descriptorpb.DescriptorProto]()
		core.AssertNotNil(t, option, "option")

		// Test that the option is created successfully by creating a resource
		res, err := resource.New[*descriptorpb.DescriptorProto](nil, option)
		core.AssertNoError(t, err, "resource creation")
		core.AssertNotNil(t, res, "resource")
	})

	t.Run("render with valid protobuf data", func(t *testing.T) {
		renderFunc := resource.RenderFunc[*descriptorpb.DescriptorProto](func(rw http.ResponseWriter, req *http.Request, data any) error {
			var msg proto.Message
			if data != nil {
				var ok bool
				msg, ok = data.(proto.Message)
				if !ok {
					return core.Wrapf(core.ErrUnreachable, "RenderProtoJSON: invalid data type: %T", data)
				}
			}
			return RenderProtoJSON(rw, req, msg)
		})

		rw := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", nil)
		testMsg := &descriptorpb.DescriptorProto{Name: proto.String("TestMessage")}

		err := renderFunc(rw, req, testMsg)
		core.AssertNoError(t, err, "render")
		core.AssertEqual(t, consts.JSON, rw.Header().Get(consts.ContentType), "content-type")
		core.AssertContains(t, rw.Body.String(), `"name":"TestMessage"`, "response body")
	})

	t.Run("render with nil protobuf data", func(t *testing.T) {
		renderFunc := resource.RenderFunc[*descriptorpb.DescriptorProto](func(rw http.ResponseWriter, req *http.Request, data any) error {
			var msg proto.Message
			if data != nil {
				var ok bool
				msg, ok = data.(proto.Message)
				if !ok {
					return core.Wrapf(core.ErrUnreachable, "RenderProtoJSON: invalid data type: %T", data)
				}
			}
			return RenderProtoJSON(rw, req, msg)
		})

		rw := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", nil)

		err := renderFunc(rw, req, nil)
		core.AssertNoError(t, err, "render")
		core.AssertEqual(t, consts.JSON, rw.Header().Get(consts.ContentType), "content-type")
		core.AssertEqual(t, "{}", rw.Body.String(), "response body")
	})

	t.Run("test WithProtoJSON option integration", func(t *testing.T) {
		// Test that WithProtoJSON can be used to create a working resource
		option := WithProtoJSON[*descriptorpb.DescriptorProto]()
		res, err := resource.New[*descriptorpb.DescriptorProto](nil, option)
		core.AssertNoError(t, err, "resource creation")
		core.AssertNotNil(t, res, "resource")

		// Verify the resource has the expected methods
		methods := res.Methods()
		core.AssertNotEqual(t, 0, len(methods), "resource should have methods")
	})
}

func TestWithJSON(t *testing.T) {
	t.Run("option creation", func(t *testing.T) {
		option := WithJSON[map[string]string]()
		core.AssertNotNil(t, option, "option")

		// Test that the option is created successfully by creating a resource
		res, err := resource.New[map[string]string](nil, option)
		core.AssertNoError(t, err, "resource creation")
		core.AssertNotNil(t, res, "resource")
	})

	t.Run("render with valid data", func(t *testing.T) {
		renderFunc := resource.RenderFunc[map[string]string](func(rw http.ResponseWriter, req *http.Request, data any) error {
			return RenderJSON(rw, req, data)
		})

		rw := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", nil)
		testData := map[string]string{"key": "value"}

		err := renderFunc(rw, req, testData)
		core.AssertNoError(t, err, "render")
		core.AssertEqual(t, consts.JSON, rw.Header().Get(consts.ContentType), "content-type")
		core.AssertEqual(t, `{"key":"value"}`, rw.Body.String(), "response body")
	})

	t.Run("render with nil data", func(t *testing.T) {
		renderFunc := resource.RenderFunc[map[string]string](func(rw http.ResponseWriter, req *http.Request, data any) error {
			return RenderJSON(rw, req, data)
		})

		rw := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", nil)

		err := renderFunc(rw, req, nil)
		core.AssertNoError(t, err, "render")
		core.AssertEqual(t, consts.JSON, rw.Header().Get(consts.ContentType), "content-type")
		core.AssertEqual(t, "null", rw.Body.String(), "response body")
	})
}

func TestHeaders(t *testing.T) {
	t.Run("content-type set correctly", func(t *testing.T) {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", nil)

		err := RenderJSON(rw, req, map[string]string{"test": "value"})
		core.AssertNoError(t, err, "render")
		core.AssertEqual(t, consts.JSON, rw.Header().Get(consts.ContentType), "content-type")
	})

	t.Run("content-length set for POST", func(t *testing.T) {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", nil)

		err := RenderJSON(rw, req, map[string]string{"test": "value"})
		core.AssertNoError(t, err, "render")

		contentLength := rw.Header().Get(consts.ContentLength)
		core.AssertNotEqual(t, "", contentLength, "content-length")
		core.AssertEqual(t, "16", contentLength, "content-length value")
	})

	t.Run("no content-length for HEAD", func(t *testing.T) {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest("HEAD", "/test", nil)

		err := RenderJSON(rw, req, map[string]string{"test": "value"})
		core.AssertNoError(t, err, "render")

		contentLength := rw.Header().Get(consts.ContentLength)
		core.AssertEqual(t, "", contentLength, "no content-length for HEAD")
	})
}

// Benchmark tests

type renderBenchmarkCase struct {
	data   any
	name   string
	method string
}

func newRenderBenchmarkCase(name, method string, data any) renderBenchmarkCase {
	return renderBenchmarkCase{
		name:   name,
		method: method,
		data:   data,
	}
}

func renderBenchmarkCases() []renderBenchmarkCase {
	smallMsg := &descriptorpb.DescriptorProto{Name: proto.String("Test")}
	largeMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("LargeMessage"),
		Field: []*descriptorpb.FieldDescriptorProto{
			{Name: proto.String("field1"), Number: proto.Int32(1)},
			{Name: proto.String("field2"), Number: proto.Int32(2)},
			{Name: proto.String("field3"), Number: proto.Int32(3)},
		},
	}

	return []renderBenchmarkCase{
		newRenderBenchmarkCase("small-proto-POST", "POST", smallMsg),
		newRenderBenchmarkCase("small-proto-HEAD", "HEAD", smallMsg),
		newRenderBenchmarkCase("large-proto-POST", "POST", largeMsg),
		newRenderBenchmarkCase("small-json-POST", "POST", map[string]string{"key": "value"}),
		newRenderBenchmarkCase("large-json-POST", "POST", map[string]any{
			"key1": "value1",
			"key2": "value2",
			"key3": map[string]string{"nested": "value"},
		}),
	}
}

func BenchmarkRenderProtoJSON(b *testing.B) {
	cases := renderBenchmarkCases()
	for _, bc := range cases {
		if _, isProto := bc.data.(proto.Message); !isProto {
			continue
		}
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rw := httptest.NewRecorder()
				req := httptest.NewRequest(bc.method, "/test", nil)
				_ = RenderProtoJSON(rw, req, bc.data.(proto.Message))
			}
		})
	}
}

func BenchmarkRenderJSON(b *testing.B) {
	cases := renderBenchmarkCases()
	for _, bc := range cases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rw := httptest.NewRecorder()
				req := httptest.NewRequest(bc.method, "/test", nil)
				_ = RenderJSON(rw, req, bc.data)
			}
		})
	}
}
