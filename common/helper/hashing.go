package helper

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the given password using bcrypt.
// It returns the hashed password as a string and any error encountered.
// Usage example:
//
//	hashedPassword, err := HashPassword("mySecretPassword")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Hashed Password:", hashedPassword)
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// ComparePassword compares a hashed password with its possible plaintext equivalent.
// It returns nil on success, or an error on failure.
// Usage example:
//
//	err := ComparePassword(hashedPassword, "mySecretPassword")
//	if err != nil {
//	    fmt.Println("Password does not match")
//	} else {
//	    fmt.Println("Password matches")
//	}
func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
