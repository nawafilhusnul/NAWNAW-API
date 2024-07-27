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
// It marshals the ID value to JSON after encrypting it using base64 encoding.
// Usage example:
//
//	id := ID(12345)
//	jsonData, err := json.Marshal(id)
//	if err != nil {
//	    fmt.Println("Error marshaling ID:", err)
//	} else {
//	    fmt.Println("Marshalled ID:", string(jsonData))
//	}
func (id ID) MarshalJSON() ([]byte, error) {
	encodedID := encryptID(id)
	return json.Marshal(encodedID)
}

// encryptID encrypts the given ID by concatenating it with a secret and encoding the result using base64.
// It first converts the ID to a string, concatenates it with the secret, and then encodes the concatenated string using base64 encoding.
// Returns the base64 encoded string.
// Usage example:
//
//	id := ID(12345)
//	encryptedID := encryptID(id)
//	fmt.Println("Encrypted ID:", encryptedID)
func encryptID(id ID) string {
	secret := vars.ENCRYPT_SECRET
	encodedID := base64.StdEncoding.EncodeToString([]byte(secret + strconv.FormatInt(int64(id), 10)))
	return encodedID
}

// decryptID decrypts a base64 encoded string to retrieve the original ID value.
// It first decodes the base64 string, removes the secret prefix, and then parses the remaining string to an int64.
// Returns the decrypted ID and an error if any occurred during the process.
// Usage example:
//
//	encryptedID := "c2VjcmV0MTIzNDU=" // base64 encoded string
//	id, err := decryptID(encryptedID)
//	if err != nil {
//	    fmt.Println("Error decrypting ID:", err)
//	} else {
//	    fmt.Println("Decrypted ID:", id)
//	}
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
// It unmarshals a JSON value into the ID struct by decrypting the base64 encoded value.
// Usage example:
//
//	var id ID
//	jsonData := `"encryptedID"`
//	err := json.Unmarshal([]byte(jsonData), &id)
//	if err != nil {
//	    fmt.Println("Error unmarshaling ID:", err)
//	} else {
//	    fmt.Println("Unmarshalled ID:", id)
//	}
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
// It scans a value from a database driver and assigns it to the ID.
// Usage example:
//
//	var id ID
//	err := id.Scan(databaseValue)
//	if err != nil {
//	    fmt.Println("Error scanning ID:", err)
//	} else {
//	    fmt.Println("Scanned ID:", id)
//	}
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
// Usage example:
//
//	id := ID(12345)
//	value, err := id.Value()
//	if err != nil {
//	    fmt.Println("Error getting value of ID:", err)
//	} else {
//	    fmt.Println("Value of ID:", value)
//	}
func (id ID) Value() (driver.Value, error) {
	if id == 0 {
		return nil, nil
	}
	return int64(id), nil
}

// ParseID parses an encrypted ID string and returns the decrypted integer value.
// Usage example:
//
//	parsedID, err := ParseID("encryptedID")
//	if err != nil {
//	    fmt.Println("Error parsing ID:", err)
//	} else {
//	    fmt.Println("Parsed ID:", parsedID)
//	}
func ParseID(id string) (int, error) {
	parsedID, err := decryptID(id)
	if err != nil {
		return 0, err
	}
	return int(parsedID), nil
}
