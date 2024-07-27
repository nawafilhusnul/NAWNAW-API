package validator

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
)

type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates the given struct based on the tags defined in the struct fields.
// It returns a responseError with a BadRequest status code if validation fails, or nil if validation succeeds.
// Usage example:
//
//	type User struct {
//	    Name  string `validate:"required"`
//	    Email string `validate:"required,email"`
//	}
//
//	user := User{Name: "John Doe", Email: "john.doe@example.com"}
//	validator := NewCustomValidator()
//	err := validator.Validate(user)
//	if err != nil {
//	    fmt.Println("Validation failed:", err)
//	} else {
//	    fmt.Println("Validation succeeded")
//	}
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return response.NewError(http.StatusBadRequest, constants.ErrorCodeBadRequest, err.Error())
	}
	return nil
}

// NewCustomValidator creates a new instance of CustomValidator with a new validator.Validate instance.
// Usage example:
//
//	validator := NewCustomValidator()
//	fmt.Println(validator)
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
