package errors

import (
	"errors"
	"fmt"
)

var (
	ErrUnknownEnvironment = errors.New("unknown environment")
	ErrE001               = errors.New("rejected")
	ErrE002               = errors.New("rejected - System error")
	ErrE003               = errors.New("invalid card")
	ErrE004               = errors.New("rejected - Invalid operation request")
	ErrE005               = errors.New("invalid merchant")
)

// Error codes
const (
	ErrorCodeE001 = "E001"
	ErrorCodeE002 = "E002"
	ErrorCodeE003 = "E003"
	ErrorCodeE004 = "E004"
	ErrorCodeE005 = "E005"
)

// UnexpectedClientResponseError is the structure that contains all relevant data of
// an error host return an unexpected response.
type UnexpectedClientResponseError struct {
	StatusCode int
	Body       []byte
}

// Error returns the string error
func (e *UnexpectedClientResponseError) Error() string {
	return fmt.Sprintf("unexpected client response: %d - %s.", e.StatusCode, string(e.Body))
}

// NewUnexpectedClientResponse returns a new UnexpectedClientResponseError concrete error
func NewUnexpectedClientResponse(statusCode int, body []byte) error {
	return &UnexpectedClientResponseError{
		StatusCode: statusCode,
		Body:       body,
	}
}

// ErrorResponse is the structure that contains all relevant data of
// an error host return an error response.
type ErrorResponse struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// Error returns the string error
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("error response: %s - %s.", e.ErrorCode, e.ErrorMessage)
}

// NewErrorResponse returns a new ErrorResponse concrete error
func NewErrorResponse(errorCode, errorMessage string) error {
	return &ErrorResponse{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}
