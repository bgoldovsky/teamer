package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/handlers/resp"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	teamsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/teams"
	"github.com/gorilla/mux"
)

const (
	defaultErrMsg       = "internal server error"
	notFoundErrMsg      = "team not found"
	invalidTeamIDErrMsg = "invalid team ID"
	notEmptyTeamErrMsg  = "team not empty"
)

type Handlers struct {
	router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
	service       *teamsSrv.Service
}

func New(router *mux.Router, signingKey string, service *teamsSrv.Service) *Handlers {
	handler := &Handlers{
		service:       service,
		router:        router,
		jwtMiddleware: middleware.NewJWT([]byte(signingKey)),
	}

	handler.jwtMiddleware.Options.ErrorHandler = resp.OnError
	handler.findRoutes()

	return handler
}

func (h *Handlers) findRoutes() {
	h.router.HandleFunc("/teams", h.GetTeams).Methods("GET")
	h.router.HandleFunc("/teams/{id}", h.GetTeam).Methods("GET")
	h.router.Handle("/teams/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.UpdateTeam))).Methods("PUT")
	h.router.Handle("/teams/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.DeleteTeam))).Methods("DELETE")
	h.router.Handle("/teams", h.jwtMiddleware.Handler(http.HandlerFunc(h.CreateTeam))).Methods("POST")
}

func (h *Handlers) GetTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, invalidTeamIDErrMsg)
		return
	}

	view, err := h.service.GetTeam(r.Context(), teamID)
	if isNotFound(err) {
		resp.RespondError(w, http.StatusNoContent, notFoundErrMsg)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	if view == nil {
		resp.RespondJSON(w, http.StatusNoContent, view)
		return
	}

	resp.RespondJSON(w, http.StatusOK, view)
}

func (h *Handlers) GetTeams(w http.ResponseWriter, r *http.Request) {
	view, err := h.service.GetTeams(r.Context())
	if isNotFound(err) {
		resp.RespondError(w, http.StatusNoContent, notFoundErrMsg)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	if len(view) == 0 {
		resp.RespondJSON(w, http.StatusNoContent, view)
		return
	}

	resp.RespondJSON(w, http.StatusOK, view)
}

func (h *Handlers) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var form models.TeamForm
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		resp.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.service.AddTeam(r.Context(), &form)
	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	resp.RespondJSON(w, http.StatusCreated, status)
}

func (h *Handlers) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, invalidTeamIDErrMsg)
		return
	}

	status, err := h.service.RemoveTeam(r.Context(), id)
	if isNotEmptyTeam(err) {
		resp.RespondError(w, http.StatusBadRequest, notEmptyTeamErrMsg)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}

func (h *Handlers) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var form models.TeamForm
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, invalidTeamIDErrMsg)
		return
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		resp.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.service.UpdateTeam(r.Context(), id, &form)
	if isNotFound(err) {
		resp.RespondError(w, http.StatusBadRequest, notFoundErrMsg)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}

func isNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), notFoundErrMsg)
}

func isNotEmptyTeam(err error) bool {
	return err != nil && strings.Contains(err.Error(), notEmptyTeamErrMsg)
}
