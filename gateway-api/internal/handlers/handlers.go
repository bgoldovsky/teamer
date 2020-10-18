package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/middleware"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	personsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/persons"
	teamsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/teams"
	"github.com/gorilla/mux"
)

const defaultErrMsg = "internal server error"

type Handlers struct {
	teamsService   *teamsSrv.Service
	personsService *personsSrv.Service
	router         *mux.Router
	jwtMiddleware  *jwtMiddleware.JWTMiddleware
}

func New(teamsService *teamsSrv.Service, personsService *personsSrv.Service) *Handlers {
	return &Handlers{
		teamsService:   teamsService,
		personsService: personsService,
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
	h.router.HandleFunc("/teams/{id}", h.GetTeam).Methods("GET")
	h.router.Handle("/teams/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.UpdateTeam))).Methods("PUT")
	h.router.Handle("/teams/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.DeleteTeam))).Methods("DELETE")
	h.router.Handle("/teams", h.jwtMiddleware.Handler(http.HandlerFunc(h.CreateTeam))).Methods("POST")

	h.router.HandleFunc("/persons", h.GetPersons).Methods("GET")
	h.router.HandleFunc("/persons/{id}", h.GetPerson).Methods("GET")
	h.router.Handle("/persons/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.UpdatePerson))).Methods("PUT")
	h.router.Handle("/persons/{id}", h.jwtMiddleware.Handler(http.HandlerFunc(h.DeletePerson))).Methods("DELETE")
	h.router.Handle("/persons", h.jwtMiddleware.Handler(http.HandlerFunc(h.CreatePerson))).Methods("POST")
}

func (h *Handlers) Run(port string) {
	logger.Log.WithField("port", port).Infoln("Server running")
	logger.Log.Fatalln(http.ListenAndServe(port, middleware.LogMiddleware(middleware.PanicMiddleware(h.router))))
}

func (h *Handlers) GetTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	view, err := h.teamsService.GetTeam(r.Context(), teamID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	if view == nil {
		respondJSON(w, http.StatusNoContent, view)
		return
	}

	respondJSON(w, http.StatusOK, view)
}

func (h *Handlers) GetTeams(w http.ResponseWriter, r *http.Request) {
	view, err := h.teamsService.GetTeams(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	if len(view) == 0 {
		respondJSON(w, http.StatusNoContent, view)
		return
	}

	respondJSON(w, http.StatusOK, view)
}

func (h *Handlers) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var form models.TeamForm
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.teamsService.AddTeam(r.Context(), &form)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusCreated, status)
}

func (h *Handlers) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	status, err := h.teamsService.RemoveTeam(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, status)
}

func (h *Handlers) UpdateTeam(w http.ResponseWriter, r *http.Request) {
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

	status, err := h.teamsService.UpdateTeam(r.Context(), id, &form)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, status)
}

func (h *Handlers) GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	view, err := h.personsService.GetPerson(r.Context(), teamID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	if view == nil {
		respondJSON(w, http.StatusNoContent, view)
		return
	}

	respondJSON(w, http.StatusOK, view)
}

func (h *Handlers) GetPersons(w http.ResponseWriter, r *http.Request) {
	view, err := h.personsService.GetPersons(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	if len(view) == 0 {
		respondJSON(w, http.StatusNoContent, view)
		return
	}

	respondJSON(w, http.StatusOK, view)
}

func (h *Handlers) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var form models.PersonForm
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status, err := h.personsService.AddPerson(r.Context(), &form)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusCreated, status)
}

func (h *Handlers) DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	status, err := h.personsService.RemovePerson(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, status)
}

func (h *Handlers) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var form models.PersonForm
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

	status, err := h.personsService.UpdatePerson(r.Context(), id, &form)
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, status)
}

func onError(w http.ResponseWriter, _ *http.Request, err string) {
	respondError(w, http.StatusUnauthorized, err)
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
