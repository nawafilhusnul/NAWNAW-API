package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type ID int64

// MarshalJSON implements the json.Marshaler interface for ID.
// It marshals the ID value to JSON after encrypting it using bcrypt.
func (id ID) MarshalJSON() ([]byte, error) {
	hashedID, err := bcrypt.GenerateFromPassword([]byte(string(rune(id))), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(hashedID))
}

// UnmarshalJSON implements the json.Unmarshaler interface for ID.
// It unmarshals a JSON value into the ID struct.
func (id *ID) UnmarshalJSON(data []byte) error {
	var hashedID string
	if err := json.Unmarshal(data, &hashedID); err != nil {
		return err
	}
	parsedID, err := strconv.ParseInt(hashedID, 10, 64)
	if err != nil {
		return err
	}
	*id = ID(parsedID)
	return nil
}

// Scan implements the Scanner interface for ID.
// It scans a value from a database driver.
func (id *ID) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	*id = ID(i.Int64)
	return nil
}

// Value implements the driver Valuer interface for ID.
// It returns the ID value or nil if the ID is zero.
func (id ID) Value() (driver.Value, error) {
	if id == 0 {
		return nil, nil
	}
	return int64(id), nil
}
