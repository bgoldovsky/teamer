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
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

const (
	msgDefaultErr    = "internal server error"
	msgTeamNotFound  = "team not found"
	msgTeamNotEmpty  = "team not empty"
	msgInvalidTeamID = "invalid team ID"
)

type Handlers struct {
	router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
	service       *teamsSrv.Service
	validate      *validator.Validate
}

func New(router *mux.Router, signingKey string, service *teamsSrv.Service) *Handlers {
	handler := &Handlers{
		service:       service,
		router:        router,
		jwtMiddleware: middleware.NewJWT([]byte(signingKey)),
		validate:      validator.New(),
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
		resp.RespondError(w, http.StatusBadRequest, msgInvalidTeamID)
		return
	}

	view, err := h.service.GetTeam(r.Context(), teamID)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isNotFound(err) {
		resp.RespondError(w, http.StatusNoContent, msgTeamNotFound)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
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
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isNotFound(err) {
		resp.RespondError(w, http.StatusNoContent, msgTeamNotFound)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
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

	err := h.validate.Struct(form)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.service.AddTeam(r.Context(), &form)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusCreated, status)
}

func (h *Handlers) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, msgInvalidTeamID)
		return
	}

	status, err := h.service.RemoveTeam(r.Context(), id)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isNotEmptyTeam(err) {
		resp.RespondError(w, http.StatusBadRequest, msgTeamNotEmpty)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}

func (h *Handlers) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var form models.TeamForm
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, msgInvalidTeamID)
		return
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		resp.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	err = h.validate.Struct(form)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.service.UpdateTeam(r.Context(), id, &form)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isNotFound(err) {
		resp.RespondError(w, http.StatusBadRequest, msgTeamNotFound)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}

func isNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), msgTeamNotFound)
}

func isNotEmptyTeam(err error) bool {
	return err != nil && strings.Contains(err.Error(), msgTeamNotEmpty)
}

func isInvalidArgument(err error) bool {
	return err != nil && strings.Contains(err.Error(), "InvalidArgument")
}
