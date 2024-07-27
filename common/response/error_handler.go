package response

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
)

// CustomHTTPErrorHandler handles HTTP errors in a custom way.
// It sets the response code and formats the error message in a custom response structure.
// Usage example:
//
//	e := echo.New()
//	e.HTTPErrorHandler = CustomHTTPErrorHandler
//	e.GET("/", func(c echo.Context) error {
//	    return echo.NewHTTPError(http.StatusBadRequest, "This is a bad request")
//	})
//	e.Start(":8080")
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
