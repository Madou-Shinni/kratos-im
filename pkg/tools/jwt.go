package tools

import (
	"github.com/golang-jwt/jwt/v4"
)

func GenToken(mp jwt.MapClaims, secret string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mp)
	signedString, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
