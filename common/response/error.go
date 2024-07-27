package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

// responseError represents an error with an associated HTTP status code, error code, and message.
type responseError struct {
	statusCode int
	errorCode  string
	message    string
}

// GetErrorStatusCode extracts the HTTP status code from the given error if it is a responseError.
// If the error is not a responseError, it returns http.StatusInternalServerError.
func GetErrorStatusCode(err error) int {
	resErr := &responseError{}
	if errors.As(err, &resErr) {
		return resErr.statusCode
	}

	return http.StatusInternalServerError
}

// Error returns the error message of the responseError.
func (e *responseError) Error() string {
	return e.message
}

// StatusCode returns the HTTP status code of the responseError.
func (e *responseError) StatusCode() int {
	return e.statusCode
}

// ErrorCode returns the error code of the responseError.
func (e *responseError) ErrorCode() string {
	return e.errorCode
}

// NewError creates a new responseError with the given HTTP status code, error code, and message.
func NewError(httpStatusCode int, errorCode string, message string) *responseError {
	return &responseError{statusCode: httpStatusCode, errorCode: errorCode, message: message}
}

// MarshalJSON customizes the JSON representation of the responseError.
func (e responseError) MarshalJSON() ([]byte, error) {
	if e.statusCode == 0 || e.errorCode == "" || e.message == "" {
		return json.Marshal(struct{}{})
	}

	return json.Marshal(struct {
		StatusCode int    `json:"status_code"`
		ErrorCode  string `json:"code"`
		Message    string `json:"message"`
	}{
		StatusCode: e.statusCode,
		ErrorCode:  e.errorCode,
		Message:    e.message,
	})
}

// UnmarshalJSON customizes the JSON deserialization of the responseError.
func (e *responseError) UnmarshalJSON(data []byte) error {
	var temp struct {
		StatusCode int    `json:"status_code"`
		ErrorCode  string `json:"code"`
		Message    string `json:"message"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	e.statusCode = temp.StatusCode
	e.errorCode = temp.ErrorCode
	e.message = temp.Message

	return nil
}
