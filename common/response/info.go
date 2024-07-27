package response

import "encoding/json"

// responseInfo represents an informational message.
type responseInfo struct {
	message string
}

// MarshalJSON customizes the JSON representation of the responseInfo.
// It returns an empty JSON object if the message is empty.
// Usage example:
//
//	info := responseInfo{message: "Operation successful"}
//	jsonData, err := info.MarshalJSON()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(string(jsonData))
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

// UnmarshalJSON customizes the JSON deserialization of the responseInfo.
// It populates the responseInfo with the data from the JSON.
// Usage example:
//
//	jsonData := []byte(`{"message": "Operation successful"}`)
//	var info responseInfo
//	err := info.UnmarshalJSON(jsonData)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(info.message)
func (r *responseInfo) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r)
}
