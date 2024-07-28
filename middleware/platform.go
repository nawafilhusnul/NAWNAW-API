package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
)

func Platform(allowedPlatforms ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(*cc.Ctx)

			if ctx.SkipCheck() {
				return next(c)
			}

			userPlatforms := ctx.GetUser().Platforms
			if v, ok := userPlatforms[constants.PlatformBasic]; !ok || !v {
				return response.NewError(http.StatusForbidden, constants.ErrorCodeForbidden, "Forbidden")
			}

			if len(allowedPlatforms) <= 0 {
				return next(c)
			}

			for _, ap := range allowedPlatforms {
				if v, ok := userPlatforms[ap]; ok && v {
					return next(c)
				}
			}

			return response.NewError(http.StatusForbidden, constants.ErrorCodeForbidden, "Forbidden")
		}
	}
}
