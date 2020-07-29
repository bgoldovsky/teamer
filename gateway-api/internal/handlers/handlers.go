package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/logger"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/middleware"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/models"
	teamsRepo "github.com/bgoldovsky/teamer-bot/gateway-api/internal/repostiory/teams"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const defaultErrMsg = "internal server error"

type Handlers struct {
	client        *teams.Client
	repo          teamsRepo.Repository
	router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
}

func New(client *teams.Client, repo teamsRepo.Repository) *Handlers {
	return &Handlers{
		client: client,
		repo:   repo,
	}
}

func (h *Handlers) Initialize(signingKey string) {
	h.jwtMiddleware = middleware.NewJWT([]byte(signingKey))
	h.jwtMiddleware.Options.ErrorHandler = onError
	h.router = mux.NewRouter()
	h.findRoutes()
}

func (h *Handlers) findRoutes() {
	h.router.HandleFunc("/teams", h.GetTeams).Methods("GET")
	h.router.Handle("/teams/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.UpdateTeam))).Methods("PUT")
	h.router.Handle("/teams/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.DeleteTeam))).Methods("DELETE")
	h.router.Handle("/teams", h.jwtMiddleware.Handler(http.HandlerFunc(h.CreateTeam))).Methods("POST")
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
		respondError(w, http.StatusInternalServerError, "token error", err)
	}
	_, _ = w.Write([]byte(tokenString))
}

func (h *Handlers) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var form models.TeamForm
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON", err)
		return
	}

	status, err := h.client.AddTeam(r.Context(), form.Name, form.Description, form.Slack)
	if err != nil {
		logger.Log.WithField("form", form).WithError(err).Errorln("add team error")
		respondError(w, http.StatusInternalServerError, defaultErrMsg, err)
		return
	}

	h.clearRepo()
	respondJSON(w, http.StatusCreated, status)
}

func (h *Handlers) GetTeams(w http.ResponseWriter, r *http.Request) {
	if view := h.getRepo(); view != nil {
		respondJSON(w, http.StatusOK, view)
		return
	}

	view, err := h.client.GetTeams(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg, err)
		return
	}

	h.saveRepo(view)

	if len(view) == 0 {
		respondJSON(w, http.StatusNoContent, view)
		return
	}

	respondJSON(w, http.StatusOK, view)
}

func (h *Handlers) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID", err)
		return
	}

	status, err := h.client.RemoveTeam(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg, err)
		return
	}

	h.clearRepo()
	respondJSON(w, http.StatusOK, status)
}

func (h *Handlers) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var form models.TeamForm
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID", err)
		return
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON", err)
		return
	}

	status, err := h.client.UpdateTeam(r.Context(), id, form.Name, form.Description, form.Slack)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg, err)
		return
	}

	h.clearRepo()
	respondJSON(w, http.StatusOK, status)
}

func onError(w http.ResponseWriter, _ *http.Request, err string) {
	respondError(w, http.StatusUnauthorized, err, errors.New(err))
}

func respondError(w http.ResponseWriter, statusCode int, msg string, err error) {
	logger.Log.Errorf("error: %s\n", err)
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

func (h *Handlers) clearRepo() {
	err := h.repo.Clear()
	if err != nil {
		logger.Log.WithError(err).Error("clear teams repo error")
	}
}

func (h *Handlers) saveRepo(teams []*models.TeamView) {
	err := h.repo.Save(teams)
	if err != nil {
		logger.Log.WithError(err).Error("save teams repo error")
	}
}

func (h *Handlers) getRepo() []*models.TeamView {
	views, err := h.repo.Get()
	if err != nil {
		logger.Log.WithError(err).Error("save teams repo error")
		return nil
	}
	return views
}
