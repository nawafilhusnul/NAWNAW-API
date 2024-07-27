package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nawafilhusnul/NAWNAW-API/common/constants"
	"github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/datatypes"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/common/vars"
)

// Session is a middleware function that checks for a valid JWT token in the Authorization header of the request.
// If the token is valid, it sets the user information in the context and calls the next handler.
// If the token is missing or invalid, it returns an unauthorized error response.
func Session() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerAuth := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")
			if len(bearerAuth) != 2 {
				return response.NewError(http.StatusUnauthorized, constants.ErrorCodeMissingToken, "Unauthorized")
			}
			tk := bearerAuth[1]

			cc, err := setInfoToCtx(c, tk)
			if err != nil {
				return err
			}

			c = cc

			return next(c)
		}
	}
}

// setInfoToCtx parses the JWT token and extracts user information from the claims.
// It sets the user information in the context and returns the updated context.
// If the token is invalid or any required claim is missing, it returns an unauthorized error response.
func setInfoToCtx(c echo.Context, tk string) (*ctx.Ctx, error) {
	claims := &jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(vars.JWT_SECRET), nil
	})
	if err != nil {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeMissingToken, "Unauthorized")
	}

	if !jwtToken.Valid {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeInvalidToken, "Unauthorized")
	}

	claims, ok := jwtToken.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeInvalidToken, "Unauthorized")
	}

	claimsUserID, ok := (*claims)["user_id"].(string)
	if !ok || claimsUserID == "" {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeInvalidToken, "Missing user_id")
	}
	userID, err := datatypes.ParseID(claimsUserID)

	if err != nil {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeInvalidToken, "Invalid user_id")
	}

	claimsRoles, ok := (*claims)["roles"].(map[string]interface{})
	if !ok {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeInvalidToken, "Missing roles")
	}

	roles := make(map[string]bool)
	for k, cp := range claimsRoles {
		if v, ok := cp.(bool); ok {
			roles[k] = v
		}
	}

	claimsPlatforms, ok := (*claims)["platforms"].(map[string]interface{})
	if !ok {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeInvalidToken, "Missing platforms")
	}

	platforms := make(map[string]bool)
	for k, cp := range claimsPlatforms {
		if v, ok := cp.(bool); ok {
			platforms[k] = v
		}
	}

	timezone, ok := (*claims)["timezone"].(string)
	if !ok || timezone == "" {
		return nil, response.NewError(http.StatusUnauthorized, constants.ErrorCodeInvalidToken, "Missing timezone")
	}

	cc := c.(*ctx.Ctx)
	cc.SetUser(&ctx.ContextUser{
		UserID:    userID,
		Roles:     roles,
		Platforms: platforms,
		Timezone:  timezone,
	})

	return cc, nil
}
