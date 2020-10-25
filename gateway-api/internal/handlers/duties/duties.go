package duties

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/handlers/resp"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	dutiesSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/duties"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

const (
	msgDefaultErr     = "internal server error"
	msgInvalidIDs     = "invalid team or persons ID"
	msgInvalidTeamID  = "invalid team ID error"
	msgPersonNotFound = "person not found"
	msgDutyNotFound   = "duty not found"
)

type Handlers struct {
	service       *dutiesSrv.Service
	router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
	validate      *validator.Validate
}

func New(router *mux.Router, signingKey string, service *dutiesSrv.Service) *Handlers {
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
	h.router.Handle("/duties/swap", h.jwtMiddleware.Handler(http.HandlerFunc(h.Swap))).Methods("PUT")
	h.router.Handle("/duties/assign", h.jwtMiddleware.Handler(http.HandlerFunc(h.Assign))).Methods("PUT")
	h.router.HandleFunc("/duties/{teamID}", h.GetCurrentDuty).Methods("GET")
	h.router.HandleFunc("/duties", h.GetDuties).Methods("GET")
}

func (h *Handlers) Swap(w http.ResponseWriter, r *http.Request) {
	var form models.SwapForm

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

	view, err := h.service.Swap(r.Context(), form.TeamId, form.FirstPersonID, form.SecondPersonID)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isInvalidIDs(err) {
		resp.RespondError(w, http.StatusBadRequest, msgInvalidIDs)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusOK, view)
}

func (h *Handlers) Assign(w http.ResponseWriter, r *http.Request) {
	var form models.AssignForm

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

	view, err := h.service.Assign(r.Context(), form.TeamId, form.PersonID)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isInvalidTeam(err) {
		resp.RespondError(w, http.StatusBadRequest, msgInvalidTeamID)
		return
	}

	if isPersonNotFound(err) {
		resp.RespondError(w, http.StatusBadRequest, msgPersonNotFound)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusOK, view)
}

func (h *Handlers) GetCurrentDuty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID, err := strconv.ParseInt(params["teamID"], 10, 64)
	if err != nil {
		resp.RespondError(w, http.StatusBadRequest, "invalid teamID")
		return
	}

	view, err := h.service.GetCurrentDuty(r.Context(), teamID)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isDutyNotFound(err) {
		resp.RespondError(w, http.StatusInternalServerError, msgDutyNotFound)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusOK, view)
}

func (h *Handlers) GetDuties(w http.ResponseWriter, r *http.Request) {
	rawTeamID := r.URL.Query().Get("team-id")
	if rawTeamID == "" {
		resp.RespondError(w, http.StatusBadRequest, "not specified teamID")
		return
	}

	teamID, err := strconv.ParseInt(rawTeamID, 10, 64)
	if err != nil || teamID == 0 {
		resp.RespondError(w, http.StatusBadRequest, "invalid teamID")
		return
	}

	rawCount := r.URL.Query().Get("count")
	if rawCount == "" {
		resp.RespondError(w, http.StatusBadRequest, "not specified count")
		return
	}

	count, err := strconv.ParseInt(rawCount, 10, 64)
	if err != nil || count == 0 {
		resp.RespondError(w, http.StatusBadRequest, "invalid count")
		return
	}

	status, err := h.service.GetDuties(r.Context(), teamID, count)
	if isInvalidArgument(err) {
		resp.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if isDutyNotFound(err) {
		resp.RespondError(w, http.StatusInternalServerError, msgDutyNotFound)
		return
	}

	if err != nil {
		resp.RespondError(w, http.StatusInternalServerError, msgDefaultErr)
		return
	}

	resp.RespondJSON(w, http.StatusOK, status)
}

func isInvalidIDs(err error) bool {
	return err != nil && strings.Contains(err.Error(), msgInvalidIDs)
}

func isInvalidTeam(err error) bool {
	return err != nil && strings.Contains(err.Error(), msgInvalidTeamID)
}

func isPersonNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), msgPersonNotFound)
}

func isDutyNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), msgDutyNotFound)
}

func isInvalidArgument(err error) bool {
	return err != nil && strings.Contains(err.Error(), "InvalidArgument")
}
