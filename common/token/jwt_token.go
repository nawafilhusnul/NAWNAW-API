package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nawafilhusnul/NAWNAW-API/common/vars"
	"github.com/nawafilhusnul/NAWNAW-API/model"
)

// generateToken generates a JWT token for a given user with a specified expiration time in seconds.
// It includes user ID, expiration time, issued at time, platforms, roles, and timezone in the token claims.
// Returns the signed token string or an error if signing fails.
func generateToken(user *model.Auth, expiredInSec int) (string, error) {
	if user.Platforms == nil {
		user.Platforms = make(map[string]bool)
	}

	if user.Roles == nil {
		user.Roles = make(map[string]bool)
	}

	if user.Permissions == nil {
		user.Permissions = make(map[string]bool)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user.ID,
		"exp":         time.Now().Add(time.Second * time.Duration(expiredInSec)).Unix(),
		"iat":         time.Now().Unix(),
		"platforms":   user.Platforms,
		"permissions": user.Permissions,
		"roles":       user.Roles,
		"timezone":    user.Timezone,
	})

	return token.SignedString([]byte(vars.JWT_SECRET))
}

// GenerateAccessToken generates an access token for a given user with the default access token expiration time.
// Returns the signed token string or an error if signing fails.
func GenerateAccessToken(user *model.Auth) (string, error) {
	return generateToken(user, vars.ACCESS_EXPIRED)
}

// GenerateRefreshToken generates a refresh token for a given user with the default refresh token expiration time.
// Returns the signed token string or an error if signing fails.
func GenerateRefreshToken(user *model.Auth) (string, error) {
	return generateToken(user, vars.REFRESH_EXPIRED)
}
