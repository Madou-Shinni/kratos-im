package common

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"kratos-im/api/errorx"
	"kratos-im/constants"
)

var (
	ErrorTokenNotFound = errorx.ErrorBus("token not found")
	ErrorParseClaims   = errorx.ErrorBus("parse claims error")
)

// GetUidFromCtx get uid from context
func GetUidFromCtx(ctx context.Context) (string, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return "", ErrorTokenNotFound
	}

	mapClaims, ok := claims.(jwtv5.MapClaims)
	if !ok {
		return "", ErrorParseClaims
	}

	return mapClaims[constants.CtxUserIDKey].(string), nil
}
