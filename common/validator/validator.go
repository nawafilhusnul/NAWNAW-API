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

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return response.NewError(http.StatusBadRequest, constants.ErrorCodeBadRequest, err.Error())
	}
	return nil
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
