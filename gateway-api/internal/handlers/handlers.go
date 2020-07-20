package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	v1 "github.com/bgoldovsky/teamer-bot/gateway-api/internal/generated/clients/people/v1"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/logger"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/middleware"
	"github.com/bgoldovsky/teamer-bot/gateway-api/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/logrusorgru/aurora"
)

const defaultErrMsg = "internal server error"

type Handlers struct {
	client        v1.TeamsClient
	Router        *mux.Router
	jwtMiddleware *jwtMiddleware.JWTMiddleware
}

func New(client v1.TeamsClient) *Handlers {
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

func (h *Handlers) GetToken(w http.ResponseWriter, r *http.Request) {
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
	var team models.Team
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&team); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	reply, err := h.client.AddTeam(r.Context(), &v1.AddTeamRequest{
		Name:        team.Name,
		Description: team.Description,
		Slack:       team.Slack,
	})

	if err != nil {
		logger.Log.WithField("team", team).WithError(err).Errorln("add team error")
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusCreated, reply.Id)
}

func (h *Handlers) getTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.client.GetTeams(r.Context(), &v1.GetTeamsRequest{
		Filter: nil,
		Limit:  1000,
		Offset: 0,
		Order:  "id",
		Sort:   "desc",
	})

	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, teams)
}

func (h *Handlers) deleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	_, err = h.client.RemoveTeam(r.Context(), &v1.RemoveTeamRequest{Id: id})
	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, "successfully removed")
}

func (h *Handlers) updateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&team); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	_, err = h.client.UpdateTeam(r.Context(), &v1.UpdateTeamRequest{
		Team: &v1.Team{
			Id:          id,
			Name:        team.Name,
			Description: team.Description,
			Slack:       team.Slack,
		},
	})

	if err != nil {
		respondError(w, http.StatusInternalServerError, defaultErrMsg)
		return
	}

	respondJSON(w, http.StatusOK, "successfully updated")
}

func respondError(w http.ResponseWriter, statusCode int, message string) {
	log.Printf("%s: %s\n", aurora.Red("Error"), message)
	respondJSON(w, statusCode, map[string]string{"error": message})
}

func respondJSON(w http.ResponseWriter, statusCode int, p interface{}) {
	resp, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)
}
