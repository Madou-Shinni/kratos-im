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
	}, time.Hour*24*7,
		"tkP2yq!i=oamTR#oQ:8n")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
}
