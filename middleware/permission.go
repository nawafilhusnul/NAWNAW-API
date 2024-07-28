package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
)

func Permission() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(*cc.Ctx)

			if ctx.SkipCheck() {
				return next(c)
			}

			userPermissions := ctx.GetUser().Permissions

			registeredRoutes := c.Echo().Routes()
			for _, rr := range registeredRoutes {
				if rr.Path == c.Request().URL.Path {
					if !userPermissions[rr.Name] {
						return response.NewError(http.StatusForbidden, constants.ErrorCodeForbidden, "Forbidden")
					}
				}
			}
			return next(c)
		}
	}
}
