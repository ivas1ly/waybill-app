package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ivas1ly/waybill-app/models"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/ivas1ly/waybill-app/database"
)

const ContextCurrentUserKey = "currentUser"

func Auth(repository database.UsersRepository) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			token, err := parseToken(request)
			if err != nil {
				handler.ServeHTTP(writer, request)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				handler.ServeHTTP(writer, request)
				return
			}

			user, err := repository.GetUserByID(claims["sub"].(string))
			if err != nil {
				handler.ServeHTTP(writer, request)
				return
			}
			ctx := context.WithValue(request.Context(), ContextCurrentUserKey, user)
			handler.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}

func getUserFromContext(ctx context.Context) (*models.User, error) {
	if ctx.Value(ContextCurrentUserKey) == nil {
		return nil, errors.New("Can't get user from context")
	}

	user, ok := ctx.Value(ContextCurrentUserKey).(*models.User)
	if !ok || user.ID == "" {
		return nil, errors.New("Can't get user from context")
	}

	return user, nil
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwt, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(viper.GetString("auth.signing_key"))
		return t, nil
	})
	return jwt, errors.Wrap(err, "JWT Parse error ")
}
