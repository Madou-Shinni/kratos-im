package tools

import (
	"github.com/golang-jwt/jwt/v4"
)

func GenToken(mp jwt.MapClaims) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mp)
	signedString, err := claims.SignedString([]byte("tkP2yq9JY2"))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
