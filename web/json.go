package web

import (
	"errors"
	"io"
	"net/http"

	// cspell:ignore protojson
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// DefaultDecodeJSONLimit is the default limit for DecodeJSON.
const DefaultDecodeJSONLimit = 1024 * 1024 // 1 MiB default

// ErrBodyTooLarge is returned when the request body is too large.
var ErrBodyTooLarge = errors.New("request body too large")

// DecodeJSON reads and parses a JSON request body into a protobuf message.
// The body is limited to the specified number of bytes for security.
//
// Parameters:
//   - req: HTTP request containing the JSON body
//   - out: protobuf message to unmarshal into
//   - limit: maximum number of bytes to read from the request body.
//     If limit <= 0, defaults to [DefaultDecodeJSONLimit].
//
// Returns an error if reading fails or JSON unmarshaling fails.
func DecodeJSON(req *http.Request, out proto.Message, limit int) error {
	if limit <= 0 {
		limit = DefaultDecodeJSONLimit
	}

	body, err := io.ReadAll(io.LimitReader(req.Body, int64(limit)+1))
	switch {
	case err != nil:
		return err
	case len(body) > limit:
		return ErrBodyTooLarge
	default:
		return protojson.Unmarshal(body, out)
	}
}
