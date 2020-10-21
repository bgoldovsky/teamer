package middleware

import (
	"net/http"
	"runtime/debug"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

func NewJWT(signingKey []byte) *jwtMiddleware.JWTMiddleware {
	return jwtMiddleware.New(jwtMiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if a := recover(); a != nil {
				logger.Log.WithFields(logrus.Fields{
					"method":    r.Method,
					"path":      r.URL.Path,
					"stack":     debug.Stack(),
					"recovered": a,
				}).Error("panic recovered")

				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("internal service error"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.
			WithField("method", r.Method).
			WithField("path", r.URL.Path).
			Infoln("request")
		next.ServeHTTP(w, r)
	})
}
