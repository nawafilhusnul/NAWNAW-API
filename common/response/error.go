package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type responseError struct {
	statusCode int
	errorCode  string
	message    string
}

func GetErrorStatusCode(err error) int {
	resErr := &responseError{}
	if errors.As(err, &resErr) {
		return resErr.statusCode
	}

	return http.StatusInternalServerError
}

func (e *responseError) Error() string {
	return e.message
}

func (e *responseError) StatusCode() int {
	return e.statusCode
}

func (e *responseError) ErrorCode() string {
	return e.errorCode
}

func NewError(httpStatusCode int, errorCode string, message string) *responseError {
	return &responseError{statusCode: httpStatusCode, errorCode: errorCode, message: message}
}

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
