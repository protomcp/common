package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"darvaza.org/core"
	"darvaza.org/x/web"
	"darvaza.org/x/web/consts"
	"darvaza.org/x/web/resource"

	// cspell:ignore protojson
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// RenderProtoJSON encodes protobuf data as JSON using protojson and sends it to the client
// after setting Content-Type and Content-Length. For HEAD only Content-Type is set.
// This ensures enums are serialized as strings rather than integers.
func RenderProtoJSON(rw http.ResponseWriter, req *http.Request, code int, data proto.Message) error {
	return doRenderJSONWithCode(rw, req, code, protojson.Marshal, data)
}

// RenderJSON encodes data as JSON using [encoding/json.Marshal] and sends it to the client
// after setting Content-Type and Content-Length. For HEAD only Content-Type is set.
func RenderJSON(rw http.ResponseWriter, req *http.Request, code int, data any) error {
	return doRenderJSONWithCode(rw, req, code, json.Marshal, data)
}

// RenderMaybeProtoJSON encodes data as JSON using protojson.Marshal if data is a protobuf message,
// otherwise it uses json.Marshal. It sends the result to the client after setting Content-Type and Content-Length.
// For HEAD only Content-Type is set.
func RenderMaybeProtoJSON(rw http.ResponseWriter, req *http.Request, code int, data any) error {
	return doRenderJSONWithCode(rw, req, code, marshalMaybeProtoJSON, data)
}

func marshalMaybeProtoJSON(data any) ([]byte, error) {
	if msg, ok := data.(proto.Message); ok {
		return protojson.Marshal(msg)
	}

	return json.Marshal(data)
}

func doRenderJSONWithCode[T any](rw http.ResponseWriter, req *http.Request, code int, fn func(T) ([]byte, error), data T) error {
	switch {
	case rw == nil, req == nil, fn == nil:
		return core.ErrInvalid
	case code < 0:
		code = http.StatusInternalServerError
	case code == 0:
		code = http.StatusOK
	}

	web.SetHeaderUnlessExists(rw.Header(), consts.ContentType, consts.JSON)

	if req.Method == consts.HEAD {
		// done
		if code == http.StatusOK {
			code = http.StatusNoContent
		}

		rw.WriteHeader(code)
		return nil
	}

	return doRenderResponse(rw, req, code, fn, data)
}

func doRenderResponse[T any](rw http.ResponseWriter, req *http.Request, code int, fn func(T) ([]byte, error), data T) error {
	b, err := fn(data)
	if err != nil {
		return err
	}

	web.SetHeaderUnlessExists(rw.Header(), consts.ContentLength, strconv.Itoa(len(b)))
	rw.WriteHeader(code)

	buf := bytes.NewBuffer(b)
	_, err = buf.WriteTo(rw)
	return err
}

// WithProtoJSON is a shortcut for WithRenderer for JSON using protojson.
// This ensures protobuf enums are serialized as strings rather than integers.
func WithProtoJSON[T proto.Message]() resource.OptionFunc[T] {
	fn := func(rw http.ResponseWriter, req *http.Request, data T) error {
		return RenderProtoJSON(rw, req, http.StatusOK, data)
	}

	return resource.WithRenderer(consts.JSON, fn)
}
