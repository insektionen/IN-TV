package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// Error is a global error that can be JSON-encoded. Wraps the real error for logging
type Error struct {
	Err     error  `json:"-"`
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Unwrap() error {
	return e.Err
}

// BadRequestError returns a new Error for 'Bad Request' with a message
func BadRequestError(message string) error {
	return Error{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

// ForbiddenError returns a new Error with Code for 'Forbidden'
func ForbiddenError() error {
	return Error{
		Code: http.StatusForbidden,
	}
}

// NotFoundError return a new error with Code for 'Not found'
func NotFoundError() error {
	return Error{
		Code: http.StatusNotFound,
	}
}

// MethodNotAllowedError return a new error with Code for 'Method not allowed'
func MethodNotAllowedError() error {
	return Error{
		Code: http.StatusMethodNotAllowed,
	}
}

// InternalServerError returns a new wrapped error with code for 'Internal Server Error'
func InternalServerError(err error) error {
	return Error{
		Err:     err,
		Message: http.StatusText(http.StatusInternalServerError),
		Code:    http.StatusInternalServerError,
	}
}

// JSONResponse encodes data as json to http.ResponseWriter
func JSONResponse(w http.ResponseWriter, data interface{}) {
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		panic(fmt.Errorf("couldn't write json response: %s", err.Error()))
	}
}

// ErrorResponse writes json error to w, also logs internal errors
func ErrorResponse(w http.ResponseWriter, err error) {
	switch err := err.(type) {
	case Error:
		switch err.Code {
		case http.StatusInternalServerError:
			if wErr := err.Unwrap(); wErr != nil {
				log.Printf("Internal Server Error: %v\n%s", wErr, string(debug.Stack()))
			} else {
				log.Printf("Internal server error without wrapped error: %s\n%s",
					err.Message, string(debug.Stack()))
			}
		case http.StatusNoContent:
			w.WriteHeader(err.Code)
			return
		}
		if err.Message == "" {
			err.Message = http.StatusText(err.Code)
		}
		w.WriteHeader(err.Code)
		JSONResponse(w, err)
	default:
		ErrorResponse(w, InternalServerError(err))
	}
}
