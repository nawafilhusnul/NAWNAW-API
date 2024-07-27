package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type HashString string

// MarshalJSON implements the json.Marshaler interface for HashString.
// It marshals the HashString value to JSON as "***".
func (hs HashString) MarshalJSON() ([]byte, error) {
	return json.Marshal("* **")
}

// UnmarshalJSON implements the json.Unmarshaler interface for HashString.
// It unmarshals a JSON value into the HashString struct.
func (hs *HashString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if str != "* * *" {
		return errors.New("invalid value for HashString")
	}
	*hs = HashString(str)
	return nil
}

// Scan implements the Scanner interface for HashString.
// It scans a value from a database driver.
func (hs *HashString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}
	*hs = HashString(s.String)
	return nil
}

// Value implements the driver Valuer interface for HashString.
// It returns the HashString value or nil if the HashString is empty.
func (hs HashString) Value() (driver.Value, error) {
	if hs == "" {
		return nil, nil
	}
	return string(hs), nil
}
