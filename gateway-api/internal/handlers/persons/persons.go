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
	personsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/persons"
	"github.com/gorilla/mux"
)

const (
	defaultErrMsg         = "internal server error"
	notFoundErrMsg        = "person not found"
	invalidTeamIDErrMsg   = "invalid team ID"
	invalidPersonIDErrMsg = "invalid person ID"
)

type Handlers struct {
	service       *personsSrv.Service
	router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
}

func New(router *mux.Router, signingKey string, service *personsSrv.Service) *Handlers {
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
	h.router.HandleFunc("/persons", h.GetPersons).Methods("GET")
	h.router.HandleFunc("/persons/{id}", h.GetPerson).Methods("GET")
	h.router.Handle("/persons/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.UpdatePerson))).Methods("PUT")
	h.router.Handle("/persons/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.DeletePerson))).Methods("DELETE")
	h.router.Handle("/persons", h.jwtMiddleware.Handler(http.HandlerFunc(h.CreatePerson))).Methods("POST")
}

func (h *Handlers) GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, invalidPersonIDErrMsg)
		return
	}

	view, err := h.service.GetPerson(r.Context(), personID)
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

func (h *Handlers) GetPersons(w http.ResponseWriter, r *http.Request) {
	var teamID *int64
	if rawTeamID := r.URL.Query().Get("teamID"); rawTeamID != "" {
		tmp, _ := strconv.ParseInt(rawTeamID, 10, 64)
		teamID = &tmp
	}

	view, err := h.service.GetPersons(r.Context(), teamID)
	if isNotFound(err) {
		resp.RespondError(w, http.StatusNoContent, defaultErrMsg)
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

func (h *Handlers) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var form models.PersonForm
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		resp.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.service.AddPerson(r.Context(), &form)
	if isInvalidTeam(err) {
		resp.RespondError(w, http.StatusBadRequest, invalidTeamIDErrMsg)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	resp.RespondJSON(w, http.StatusCreated, status)
}

func (h *Handlers) DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, invalidPersonIDErrMsg)
		return
	}

	status, err := h.service.RemovePerson(r.Context(), id)
	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}

func (h *Handlers) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var form models.PersonForm
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, invalidPersonIDErrMsg)
		return
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		resp.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.service.UpdatePerson(r.Context(), id, &form)
	if isNotFound(err) {
		resp.RespondError(w, http.StatusBadRequest, notFoundErrMsg)
		return
	}

	if isInvalidTeam(err) {
		resp.RespondError(w, http.StatusBadRequest, invalidTeamIDErrMsg)
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

func isInvalidTeam(err error) bool {
	return err != nil && strings.Contains(err.Error(), invalidTeamIDErrMsg)
}
