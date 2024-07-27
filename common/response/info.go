package response

import "encoding/json"

type responseInfo struct {
	message string
}

func (r responseInfo) MarshalJSON() ([]byte, error) {
	if r.message == "" {
		return json.Marshal(struct{}{})
	}

	return json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: r.message,
	})
}

func (r *responseInfo) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r)
}
