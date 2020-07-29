package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bgoldovsky/teamer/auth-api/internal/logger"
	"github.com/bgoldovsky/teamer/auth-api/internal/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const defaultErrMsg = "internal server error"

type Handlers struct {
	router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
}

func New() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Initialize(signingKey string) {
	h.jwtMiddleware = middleware.NewJWT([]byte(signingKey))
	h.router = mux.NewRouter()
	h.findRoutes()
}

func (h *Handlers) findRoutes() {
	h.router.HandleFunc("/token", h.GetToken).Methods("GET")
}

func (h *Handlers) Run(port string) {
	logger.Log.WithField("port", port).Infoln("Server running")
	logger.Log.Fatalln(http.ListenAndServe(port, middleware.LogMiddleware(middleware.PanicMiddleware(h.router))))
}

func (h *Handlers) GetToken(w http.ResponseWriter, _ *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["admin"] = true
	claims["name"] = "JohnDoe"

	token.Claims = claims
	signString, _ := h.jwtMiddleware.Options.ValidationKeyGetter(token)
	tokenString, err := token.SignedString(signString)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
	}
	_, _ = w.Write([]byte(tokenString))
}

func respondError(w http.ResponseWriter, statusCode int, msg string) {
	respondJSON(w, statusCode, map[string]string{"error": msg})
}

func respondJSON(w http.ResponseWriter, statusCode int, p interface{}) {
	resp, err := json.Marshal(p)
	if err != nil {
		logger.Log.WithError(err).Error("serialization error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)
}
