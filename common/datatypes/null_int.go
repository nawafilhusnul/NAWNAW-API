package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullInt struct {
	Int   int
	Valid bool
}

// Scan implements the Scanner interface for NullInt.
// It scans a value from a database driver.
func (ni *NullInt) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	ni.Int, ni.Valid = int(i.Int64), i.Valid
	return nil
}

// Value implements the driver Valuer interface for NullInt.
// It returns the int value or nil if the int is not valid.
func (ni NullInt) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return int64(ni.Int), nil
}

// MarshalJSON implements the json.Marshaler interface for NullInt.
// It marshals the int value to JSON or null if the int is not valid.
func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(ni.Int)
}

// UnmarshalJSON implements the json.Unmarshaler interface for NullInt.
// It unmarshals a JSON value into the NullInt struct.
func (ni *NullInt) UnmarshalJSON(data []byte) error {
	var i *int
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i != nil {
		ni.Int = *i
		ni.Valid = *i != 0
	} else {
		ni.Valid = false
	}

	return nil
}

// SetNullInt sets the int and valid flag for a NullInt struct.
// If the int is not the zero value, the valid flag is set to true.
// If the valid flag is not set, it defaults to true if the int is not the zero value.
func SetNullInt(i int, valid ...bool) NullInt {
	v := i != 0
	if len(valid) > 0 {
		v = valid[0]
	}
	return NullInt{Int: i, Valid: v}
}
