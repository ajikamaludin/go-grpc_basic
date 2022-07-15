package errors

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Response struct {
	Success bool   `json:"success,omitempty"`
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
}

const (
	FallbackError       = `{"status": false, "code": "100", "msg": "internal error"}`
	InternalServerError = `Internal Server Error`
)

// CustomHTTPError is the custom http error handler for grpc gateway.
func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, ierr error) {
	var desc string

	// set header
	w.Header().Set("Content-Type", marshaler.ContentType())

	if s, ok := status.FromError(ierr); ok {
		// set http status code
		w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))

		desc = s.Message()
	} else {
		w.WriteHeader(runtime.HTTPStatusFromCode(codes.Unknown))
		desc = ierr.Error()
	}
	b := new(Response)
	err := json.Unmarshal([]byte(desc), b)
	if err != nil {
		_, _ = w.Write([]byte(FallbackError))
	} else {
		err = json.NewEncoder(w).Encode(b)
		if err != nil {
			_, _ = w.Write([]byte(FallbackError))
		}
	}
}

// FormatError is the exposed function for generating errors.
func FormatError(c codes.Code, message *Response) error {
	if message == nil {
		return status.Errorf(c, codes.Internal.String())
	}

	buf, err := json.Marshal(message)
	if err != nil {
		return status.Errorf(c, codes.Internal.String())
	}

	// TODO: logger store error

	return status.Errorf(c, string(buf))
}

// func to create error from string
func New(errs string) error {
	return errors.New(errs)
}
