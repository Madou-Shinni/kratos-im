package handle

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

const (
	UserIDCtxKey     = "userId"
	authorizationKey = "Authorization"
	bearerWord       = "Bearer"
	reason           = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

type authKey struct{}

// Option is jwt option.
type Option func(*options)

// Parser is a jwt parser
type options struct {
	signingMethod jwt.SigningMethod
	claims        func() jwt.Claims
	tokenHeader   map[string]interface{}
}

// WithSigningMethod with signing method option.
func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// WithClaims with customer claim
// If you use it in Server, f needs to return a new jwt.Claims object each time to avoid concurrent write problems
// If you use it in Client, f only needs to return a single object to provide performance
func WithClaims(f func() jwt.Claims) Option {
	return func(o *options) {
		o.claims = f
	}
}

// WithTokenHeader withe customer tokenHeader for client side
func WithTokenHeader(header map[string]interface{}) Option {
	return func(o *options) {
		o.tokenHeader = header
	}
}

type JWTAuth struct {
	keyFunc jwt.Keyfunc
	options *options
}

func NewJWTAuth(keyFunc jwt.Keyfunc, opts ...Option) *JWTAuth {
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
	}
	for _, opt := range opts {
		opt(o)
	}
	return &JWTAuth{
		keyFunc: keyFunc,
		options: o,
	}
}

func (j *JWTAuth) Auth(w http.ResponseWriter, r *http.Request) (bool, error) {
	auths := strings.SplitN(r.Header.Get(authorizationKey), " ", 2)
	if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
		return false, ErrMissingJwtToken
	}
	jwtToken := auths[1]
	var (
		tokenInfo *jwt.Token
		err       error
	)
	if j.options.claims != nil {
		tokenInfo, err = jwt.ParseWithClaims(jwtToken, j.options.claims(), j.keyFunc)
	} else {
		tokenInfo, err = jwt.Parse(jwtToken, j.keyFunc)
	}
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return false, errors.Unauthorized(reason, err.Error())
		}
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, ErrTokenInvalid
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, ErrTokenExpired
		}
		if ve.Inner != nil {
			return false, ve.Inner
		}
		return false, ErrTokenParseFail
	}
	if !tokenInfo.Valid {
		return false, ErrTokenInvalid
	}
	if tokenInfo.Method != j.options.signingMethod {
		return false, ErrUnSupportSigningMethod
	}

	*r = *r.WithContext(context.WithValue(r.Context(), authKey{}, tokenInfo.Claims))
	return true, nil
}

func (j *JWTAuth) UserId(r *http.Request) string {
	token, ok := r.Context().Value(authKey{}).(jwt.Claims)
	if !ok {
		return ""
	}

	claims := token.(jwt.MapClaims)
	if userID, ok := claims[UserIDCtxKey]; ok {
		return userID.(string)
	}

	return ""
}
