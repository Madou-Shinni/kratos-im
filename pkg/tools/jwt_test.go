package tools

import (
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(jwt.MapClaims{
		"userId": "9aec6f89-c1b6-4d85-b64d-66cf211a007f",
	}, time.Hour*24*30,
		"tkP2yq!i=oamTR#oQ:8n")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
}
