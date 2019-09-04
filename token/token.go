package token

import "context"
import "net/http"
import "strings"
import "errors"
import "encoding/base64"
import "encoding/json"

// import kithttp "github.com/go-kit/kit/transport/http"

const (
	AuthHeader = "Authorization"
)

type Languages []string

type TokenPayload struct {
	Username       string    `json:"username"`
	MainLanguage   string    `json:"mainLanguage"`
	MainLanguageId int       `json:"mainLanguageId"`
	Languages      Languages `json:"languages"`
	Role           string    `json:"role"`
	Hash           string    `json:"hash"`
	Uid            int       `json:"uid"`
	Gid            int       `json:"gid"`
}

func HttpServerBeforeTokenContextMiddleware(ctx context.Context, r *http.Request) context.Context {
	token := r.Header.Get(AuthHeader)

	ctx = context.WithValue(ctx, AuthHeader, token)
	return ctx
}

// func DecodeRequestFuncMiddleware(next kithttp.DecodeRequestFunc) kithttp.DecodeRequestFunc {
// 	// ignore token validation
// 	if true == true {
// 		return next
// 	}
// 	return func(ctx context.Context, request *http.Request) (interface{}, error) {
// 		if _, err := DecodedTokenPayload(ctx); err != nil {
// 			return nil, err
// 		}
// 		return next(ctx, request)
// 	}
// }

func HttpSetTokenInHeaderForRoutingMiddleware(ctx context.Context, w *http.Request) context.Context {
	token := ""
	if ctx.Value(AuthHeader) != nil {
		token = ctx.Value(AuthHeader).(string)
	}

	if token != "" {
		w.Header.Set(AuthHeader, token)
	}

	return ctx
}

func GetToken(ctx context.Context) (string, error) {
	tokenString := ""
	if ctx.Value(AuthHeader) != nil {
		tokenString = ctx.Value(AuthHeader).(string)
	}

	if tokenString == "" {
		return "", ErrMissingToken
	}

	splited := strings.Fields(tokenString)
	if len(splited) != 2 {

		return "", ErrWrongFormatedBearer
	}
	if len(splited) == 2 && splited[0] == "Bearer" {
		// return false, ErrInvalidToken

		tokenString = splited[1]
	} else {
		return "", ErrWrongFormatedBearer
	}

	if tokenString == "" {
		return "", ErrMissingToken
	}

	return tokenString, nil
}

func DecodedTokenPayload(ctx context.Context) (*TokenPayload, error) {
	tokenString, err := GetToken(ctx)
	if err != nil {
		return nil, err
	}

	splited := strings.Split(tokenString, ".")
	if len(splited) != 3 {
		return nil, ErrWrongFormatedToken
	}

	jsonPayload, err := base64.RawStdEncoding.DecodeString(splited[1])

	if err != nil {
		return nil, err
	}

	claims := &TokenPayload{}
	err = json.Unmarshal([]byte(jsonPayload), claims)

	return claims, err
}

var ErrInvalidToken = errors.New("invalid token")
var ErrMissingToken = errors.New("missing token")
var ErrWrongFormatedBearer = errors.New("wrong formated bearer")
var ErrWrongFormatedToken = errors.New("wrong formated token")
