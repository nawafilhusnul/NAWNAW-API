package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"strconv"

	"github.com/nawafilhusnul/NAWNAW-API/common/vars"
)

type ID int64

// MarshalJSON implements the json.Marshaler interface for ID.
// It marshals the ID value to JSON after encrypting it using bcrypt.
func (id ID) MarshalJSON() ([]byte, error) {
	encodedID := encryptID(id)
	return json.Marshal(encodedID)
}

func encryptID(id ID) string {
	secret := vars.ENCRYPT_SECRET
	encodedID := base64.StdEncoding.EncodeToString([]byte(secret + strconv.FormatInt(int64(id), 10)))
	return encodedID
}

func decryptID(hashedID string) (ID, error) {
	decodedID, err := base64.StdEncoding.DecodeString(hashedID)
	if err != nil {
		return 0, err
	}
	secret := vars.ENCRYPT_SECRET
	parsedID, err := strconv.ParseInt(string(decodedID[len(secret):]), 10, 64)
	if err != nil {
		return 0, err
	}
	return ID(parsedID), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ID.
// It unmarshals a JSON value into the ID struct.
func (id *ID) UnmarshalJSON(data []byte) error {
	var hashedID string
	if err := json.Unmarshal(data, &hashedID); err != nil {
		return err
	}

	parsedID, err := decryptID(hashedID)
	if err != nil {
		return err
	}

	*id = parsedID
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

func ParseID(id string) (int, error) {
	parsedID, err := decryptID(id)
	if err != nil {
		return 0, err
	}
	return int(parsedID), nil
}
