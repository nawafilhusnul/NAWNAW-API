package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullString struct {
	String string
	Valid  bool
}

// Scan implements the Scanner interface for NullString.
// It scans a value from a database driver.
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}
	ns.String, ns.Valid = s.String, s.Valid
	return nil
}

// Value implements the driver Valuer interface for NullString.
// It returns the string value or nil if the string is not valid.
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

// MarshalJSON implements the json.Marshaler interface for NullString.
// It marshals the string value to JSON or null if the string is not valid.
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON implements the json.Unmarshaler interface for NullString.
// It unmarshals a JSON value into the NullString struct.
func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		ns.String = *s
		ns.Valid = true
	} else {
		ns.Valid = false
	}

	return nil
}

// SetNullString sets the string and valid flag for a NullString struct.
// If the string is not the empty string, the valid flag is set to true.
// If the valid flag is not set, it defaults to true if the string is not empty.
func SetNullString(s string, valid ...bool) NullString {
	v := s != ""
	if len(valid) > 0 {
		v = valid[0]
	}
	return NullString{String: s, Valid: v}
}
