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
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

const (
	msgDefaultErr      = "internal server error"
	msgPersonNotFound  = "person not found"
	msgInvalidTeam     = "invalid team ID"
	msgInvalidPersonID = "invalid person ID"
)

type Handlers struct {
	service       *personsSrv.Service
	router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
	validate      *validator.Validate
}

func New(router *mux.Router, signingKey string, service *personsSrv.Service) *Handlers {
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
		resp.RespondError(w, http.StatusBadRequest, msgInvalidPersonID)
		return
	}

	view, err := h.service.GetPerson(r.Context(), personID)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isNotFound(err) {
		resp.RespondError(w, http.StatusNoContent, msgPersonNotFound)
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

func (h *Handlers) GetPersons(w http.ResponseWriter, r *http.Request) {
	var teamID *int64
	if rawTeamID := r.URL.Query().Get("teamID"); rawTeamID != "" {
		tmp, _ := strconv.ParseInt(rawTeamID, 10, 64)
		teamID = &tmp
	}

	view, err := h.service.GetPersons(r.Context(), teamID)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isNotFound(err) {
		resp.RespondError(w, http.StatusNoContent, msgDefaultErr)
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

func (h *Handlers) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var form models.PersonForm
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

	status, err := h.service.AddPerson(r.Context(), &form)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isInvalidTeam(err) {
		resp.RespondError(w, http.StatusBadRequest, msgInvalidTeam)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusCreated, status)
}

func (h *Handlers) DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, msgInvalidPersonID)
		return
	}

	status, err := h.service.RemovePerson(r.Context(), id)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}

func (h *Handlers) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var form models.PersonForm
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, msgInvalidPersonID)
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

	status, err := h.service.UpdatePerson(r.Context(), id, &form)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isNotFound(err) {
		resp.RespondError(w, http.StatusBadRequest, msgPersonNotFound)
		return
	}

	if isInvalidTeam(err) {
		resp.RespondError(w, http.StatusBadRequest, msgInvalidTeam)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}

func isNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), msgPersonNotFound)
}

func isInvalidTeam(err error) bool {
	return err != nil && strings.Contains(err.Error(), msgInvalidTeam)
}

func isInvalidArgument(err error) bool {
	return err != nil && strings.Contains(err.Error(), "InvalidArgument")
}
