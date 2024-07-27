package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullBool struct {
	Bool  bool
	Valid bool
}

// Scan implements the Scanner interface for NullBool.
// It scans a value from a database driver.
func (nb *NullBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}
	nb.Bool, nb.Valid = b.Bool, b.Valid
	return nil
}

// Value implements the driver Valuer interface for NullBool.
// It returns the bool value or nil if the bool is not valid.
func (nb NullBool) Value() (driver.Value, error) {
	if !nb.Valid {
		return nil, nil
	}
	return nb.Bool, nil
}

// MarshalJSON implements the json.Marshaler interface for NullBool.
// It marshals the bool value to JSON or null if the bool is not valid.
func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON implements the json.Unmarshaler interface for NullBool.
// It unmarshals a JSON value into the NullBool struct.
func (nb *NullBool) UnmarshalJSON(data []byte) error {
	var b *bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		nb.Bool = *b
		nb.Valid = true
	} else {
		nb.Valid = false
	}

	return nil
}

// SetNullBool sets the bool and valid flag for a NullBool struct.
// If the bool is not the zero value, the valid flag is set to true.
// If the valid flag is not set, it defaults to true if the bool is not the zero value.
func SetNullBool(b bool, valid ...bool) NullBool {
	v := b
	if len(valid) > 0 {
		v = valid[0]
	}
	return NullBool{Bool: b, Valid: v}
}
