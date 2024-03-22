package rws

import (
	"fmt"
	"net/http"
	"time"
)

// Authentication is an interface for authentication.
type Authentication interface {
	Auth(w http.ResponseWriter, r *http.Request) (bool, error)
	UserId(R *http.Request) string
}

type DefaultAuthentication struct {
}

func (a *DefaultAuthentication) Auth(w http.ResponseWriter, r *http.Request) (bool, error) {
	return true, nil
}

func (a *DefaultAuthentication) UserId(r *http.Request) string {
	query := r.URL.Query()
	if userID := query.Get("userId"); userID != "" {
		return userID
	}
	return fmt.Sprint(time.Now().UnixMilli())
}
