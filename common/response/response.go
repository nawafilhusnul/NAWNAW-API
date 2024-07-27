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

// NewResponse creates a new response instance.
// Usage example:
//
//	resp := NewResponse()
//	fmt.Println(resp)
func NewResponse() *response {
	return &response{}
}

// WithError sets the error information in the response.
// If the provided error is nil, it returns the response as is.
// If the error is of type responseError, it sets it in the meta field.
// Otherwise, it creates a new responseError with a BadRequest status code and sets it in the meta field.
// Usage example:
//
//	resp := NewResponse().WithError(errors.New("an error occurred"))
//	fmt.Println(resp)
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

// WithData sets the data and informational message in the response.
// If there is already an error set in the meta field, it returns the response as is.
// Usage example:
//
//	resp := NewResponse().WithData(map[string]string{"key": "value"}, "Operation successful")
//	fmt.Println(resp)
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

// MarshalJSON customizes the JSON representation of the response.
// Usage example:
//
//	resp := NewResponse().WithData(map[string]string{"key": "value"}, "Operation successful")
//	jsonData, err := resp.MarshalJSON()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(string(jsonData))
func (r response) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Meta meta `json:"meta"`
		Data any  `json:"data"`
	}{
		Meta: r.meta,
		Data: r.data,
	})
}

// UnmarshalJSON customizes the JSON deserialization of the response.
// It populates the response with the data from the JSON.
// Usage example:
//
//	jsonData := []byte(`{"meta": {"info": {"message": "Operation successful"}}, "data": {"key": "value"}}`)
//	var resp response
//	err := resp.UnmarshalJSON(jsonData)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(resp)
func (r *response) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

// MarshalJSON customizes the JSON representation of the meta.
// Usage example:
//
//	m := meta{info: responseInfo{message: "Operation successful"}}
//	jsonData, err := m.MarshalJSON()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(string(jsonData))
func (m meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Error responseError `json:"error"`
		Info  responseInfo  `json:"info"`
	}{
		Error: m.error,
		Info:  m.info,
	})
}

// UnmarshalJSON customizes the JSON deserialization of the meta.
// It populates the meta with the data from the JSON.
// Usage example:
//
//	jsonData := []byte(`{"error": {"status_code": 400, "code": "bad_request", "message": "An error occurred"}, "info": {"message": "Operation successful"}}`)
//	var m meta
//	err := m.UnmarshalJSON(jsonData)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(m)
func (m *meta) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, m)
}
