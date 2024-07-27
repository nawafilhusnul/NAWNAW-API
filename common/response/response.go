package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
)

type response struct {
	meta meta `json:"meta"`
	data any  `json:"data"`
}

type meta struct {
	error responseError `json:"error,omitempty"`
	info  responseInfo  `json:"info,omitempty"`
}

func NewResponse() *response {
	return &response{}
}

func (r *response) WithError(err error) *response {
	if err == nil {
		return r
	}

	resErr := &responseError{}
	if errors.As(err, &resErr) {
		r.meta.error = *resErr
		return r
	}

	r.meta.error = responseError{
		statusCode: http.StatusBadRequest,
		errorCode:  constants.ErrorCodeBadRequest,
		message:    err.Error(),
	}

	return r
}

func (r *response) WithData(data any, message string) *response {
	if r.meta.error != (responseError{}) {
		return r
	}

	r.data = data
	r.meta.info = responseInfo{
		message: message,
	}

	return r
}

func (r response) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Meta meta `json:"meta"`
		Data any  `json:"data"`
	}{
		Meta: r.meta,
		Data: r.data,
	})
}

func (r *response) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

func (m meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Error responseError `json:"error"`
		Info  responseInfo  `json:"info"`
	}{
		Error: m.error,
		Info:  m.info,
	})
}

func (m *meta) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, m)
}
