package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nawafilhusnul/NAWNAW-API/common/vars"
	"github.com/nawafilhusnul/NAWNAW-API/model"
)

func generateToken(user *model.Auth, expiredInSec int) (string, error) {
	if user.Platforms == nil {
		user.Platforms = make(map[string]bool)
	}

	if user.Roles == nil {
		user.Roles = make(map[string]bool)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"exp":       time.Now().Add(time.Second * time.Duration(expiredInSec)).Unix(),
		"iat":       time.Now().Unix(),
		"platforms": user.Platforms,
		"roles":     user.Roles,
		"timezone":  user.Timezone,
	})

	return token.SignedString([]byte(vars.JWT_SECRET))
}

func GenerateAccessToken(user *model.Auth) (string, error) {
	return generateToken(user, vars.ACCESS_EXPIRED)
}

func GenerateRefreshToken(user *model.Auth) (string, error) {
	return generateToken(user, vars.REFRESH_EXPIRED)
}
