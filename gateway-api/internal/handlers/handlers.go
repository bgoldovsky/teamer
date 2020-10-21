package handlers

import (
	"net/http"

	personHandlers "github.com/bgoldovsky/dutyer/gateway-api/internal/handlers/persons"
	teamHandlers "github.com/bgoldovsky/dutyer/gateway-api/internal/handlers/teams"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/middleware"
	personsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/persons"
	teamsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/teams"
	"github.com/gorilla/mux"
)

type Handler struct {
	router *mux.Router
}

func NewHandler(
	teamsService *teamsSrv.Service,
	personsService *personsSrv.Service,
	signingKey string,
) *Handler {
	router := mux.NewRouter()
	teamHandlers.New(router, signingKey, teamsService)
	personHandlers.New(router, signingKey, personsService)

	return &Handler{
		router: router,
	}
}

func (h *Handler) Run(port string) {
	logger.Log.WithField("port", port).Infoln("Server running")
	logger.Log.Fatalln(http.ListenAndServe(port, middleware.LogMiddleware(middleware.PanicMiddleware(h.router))))
}
