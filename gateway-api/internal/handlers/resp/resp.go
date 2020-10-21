package resp

import (
	"encoding/json"
	"net/http"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
)

func OnError(w http.ResponseWriter, _ *http.Request, err string) {
	RespondError(w, http.StatusUnauthorized, err)
}

func RespondError(w http.ResponseWriter, statusCode int, msg string) {
	RespondJSON(w, statusCode, map[string]string{"error": msg})
}

func RespondJSON(w http.ResponseWriter, statusCode int, p interface{}) {
	resp, err := json.Marshal(p)
	if err != nil {
		logger.Log.WithError(err).Error("serialization error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)
}
