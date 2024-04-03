package tools

import (
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(jwt.MapClaims{
		"userId": "6736fd5e-cac4-4d33-863b-78904382ad96",
	}, time.Hour*24*30,
		"tkP2yq!i=oamTR#oQ:8n")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
}
