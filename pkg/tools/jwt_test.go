package tools

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(jwt.MapClaims{
		"userId": uuid.New().String(),
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
}
