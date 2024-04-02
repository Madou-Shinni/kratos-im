package tools

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenToken(mp jwt.MapClaims, exp time.Duration, secret string) (string, error) {
	mp["exp"] = time.Now().Add(exp).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mp)
	signedString, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
