package middleware

import (
	"strings"
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	
	"github.com/dickywijayaa/orders-go-graphql/repositories"
	"github.com/dickywijayaa/orders-go-graphql/models"
)

const JWT_SECRET = "jwtsecret"
const CURRENT_USER_KEY = "currentuser"

func OAuthMiddleware(repo repositories.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := repo.GetUserById(claims["jti"].(string))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CURRENT_USER_KEY, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter: removePrefixFromToken,
}

func removePrefixFromToken(token string) (string, error) {
	prefix := "BEARER"

	if len(token) > len(prefix) && strings.ToUpper(token[0:len(prefix)]) == prefix {
		return token[len(prefix)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(JWT_SECRET)
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

func GetCurrentUserFromCTX(ctx context.Context) (*models.User, error) {
	errNoUserInContext := errors.New("no user in context")

	if ctx.Value(CURRENT_USER_KEY) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(CURRENT_USER_KEY).(*models.User)
	if !ok || user.ID == "" {
		return nil, errNoUserInContext
	}

	return user, nil
}