package response

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	customError := &responseError{}
	if errors.As(err, &customError) {
		c.JSON(code, NewResponse().WithError(customError))
		return
	}

	err = NewError(code, constants.ErrorCodeInternalServerError, err.Error())
	c.JSON(code, NewResponse().WithError(err))
}
