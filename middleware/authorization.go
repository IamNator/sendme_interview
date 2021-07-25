package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/IamNator/sendme_interview/internal/schema"
	"github.com/IamNator/sendme_interview/utils/httperror"
)

//ContextKey is used to access data passed by context to request
var ContextKey = schema.User{}

//retrieveBearerToken ... extracts the bearer token from request object
func retrieveBeareToken(r *http.Request) (string, *httperror.Error) {

	bearerToken := r.Header.Get("Authorization")
	auth := strings.Split(bearerToken, " ") // BEARER TOKEN: faskfn452knfdk
	if len(auth) < 2 {
		return "", httperror.New2(http.StatusUnauthorized, errors.New("unauthorized access, bearer token missing"), nil)
	}

	token := strings.TrimSpace(auth[1]) //trim excess space
	if token == "" {
		return "", httperror.New2(http.StatusUnauthorized, errors.New("unauthorized access"), nil)

	}

	return token, nil
}

//ValidateToken ....
func ValidateToken(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token, errr := retrieveBeareToken(r)
			if errr != nil {
				httperror.Default(errr).Reply(w)
				return
			}

			var user schema.User
			result := db.Table(schema.User{}.TableName()).Where("token = ?", token).
				First(&user)

			if result.RecordNotFound() {
				httperror.Default(fmt.Errorf("Unauthorized Access -")).ReplyUnauthorizedResponse(w)
				return
			}

			ttmp, er := time.Parse("2006-01-02 3:04PM", user.TokenExpiration)
			if er != nil {

				httperror.Default(er).ReplyInternalServerError(w)
				return
			}

			if ttmp.Before(time.Now()) {

				httperror.Default(fmt.Errorf("Unauthorized Access")).ReplyUnauthorizedResponse(w)
				return
			}

			ctx := context.WithValue(r.Context(), ContextKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})

	}
}
