package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/handlers/resp"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	personsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/persons"
	"github.com/gorilla/mux"
)

const defaultErrMsg = "internal server error"

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
	teamID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	view, err := h.service.GetPerson(r.Context(), teamID)
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
	view, err := h.service.GetPersons(r.Context())
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
		resp.RespondError(w, http.StatusBadRequest, "invalid ID")
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
		resp.RespondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		resp.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.service.UpdatePerson(r.Context(), id, &form)
	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}
