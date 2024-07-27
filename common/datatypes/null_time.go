package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

// Scan implements the Scanner interface for NullTime.
// It scans a value from a database driver and assigns it to the NullTime struct.
// If the scanned value is null, the Valid field is set to false.
func (nt *NullTime) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}
	nt.Time, nt.Valid = t.Time, t.Valid
	return nil
}

// Value implements the driver Valuer interface for NullTime.
// It returns the time value or nil if the time is not valid.
// This method is used to convert the NullTime struct to a driver.Value
// that can be stored in a database.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// MarshalJSON implements the json.Marshaler interface for NullTime.
// It marshals the time value to JSON or null if the time is not valid.
// This method is used to convert the NullTime struct to a JSON representation.
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nt.Time)
}

// UnmarshalJSON implements the json.Unmarshaler interface for NullTime.
// It unmarshals a JSON value into the NullTime struct.
// If the JSON value is null, the Valid field is set to false.
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	var t *time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	if t != nil {
		nt.Time = *t
		nt.Valid = true
	} else {
		nt.Valid = false
	}

	return nil
}

// SetNullTime sets the time and valid flag for a NullTime struct.
// If the time is not the zero value, the valid flag is set to true.
// If the valid flag is not set, it defaults to true if the time is not the zero value.
// Usage example:
//
//	nt := SetNullTime(time.Now())
//	nt := SetNullTime(time.Time{}, false)
func SetNullTime(t time.Time, valid ...bool) NullTime {
	v := !t.IsZero()
	if len(valid) > 0 {
		v = valid[0]
	}
	return NullTime{Time: t, Valid: v}
}
