package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/clients/teams"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/logger"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/middleware"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/logrusorgru/aurora"
)

const defaultErrMsg = "internal server error"

type Handlers struct {
	client        *teams.Client
	Router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
}

func New(client *teams.Client) *Handlers {
	return &Handlers{
		client: client,
	}
}

func (h *Handlers) Initialize(signingKey string) {
	h.jwtMiddleware = middleware.NewJWT([]byte(signingKey))
	h.Router = mux.NewRouter()
	h.findRoutes()
}

func (h *Handlers) findRoutes() {
	h.Router.HandleFunc("/teams", h.getTeams).Methods("GET")
	h.Router.Handle("/teams/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.updateTeam))).Methods("PUT")
	h.Router.Handle("/teams/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.deleteTeam))).Methods("DELETE")
	h.Router.Handle("/teams", h.jwtMiddleware.Handler(http.HandlerFunc(h.createTeam))).Methods("POST")
	h.Router.HandleFunc("/token", h.GetToken).Methods("GET")
}

func (h *Handlers) Run(port string) {
	logger.Log.WithField("port", port).Infoln("Server running")
	logger.Log.Fatalln(http.ListenAndServe(port, middleware.LogMiddleware(middleware.PanicMiddleware(h.Router))))
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
		respondError(w, http.StatusInternalServerError, "token error")
	}
	_, _ = w.Write([]byte(tokenString))
}

func (h *Handlers) createTeam(w http.ResponseWriter, r *http.Request) {
	var form models.TeamForm
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.client.AddTeam(r.Context(), form.Name, form.Description, form.Slack)

	if err != nil {
		logger.Log.WithField("form", form).WithError(err).Errorln("add team error")
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusCreated, status)
}

func (h *Handlers) getTeams(w http.ResponseWriter, r *http.Request) {
	view, err := h.client.GetTeams(r.Context())

	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, view)
}

func (h *Handlers) deleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	status, err := h.client.RemoveTeam(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, status)
}

func (h *Handlers) updateTeam(w http.ResponseWriter, r *http.Request) {
	var form models.TeamForm
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.client.UpdateTeam(r.Context(), id, form.Name, form.Description, form.Slack)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, status)
}

func respondError(w http.ResponseWriter, statusCode int, message string) {
	log.Printf("%s: %s\n", aurora.Red("Error"), message)
	respondJSON(w, statusCode, map[string]string{"error": message})
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
