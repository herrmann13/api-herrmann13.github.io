package system

import (
	"net/http"
	"encoding/json"
	"time"	
)

type HealthResponse struct {
	Status string `json:"status"`
	Service string `json:"service"`
	UptimeSeconds int64 `json:"uptime_seconds"`
	Timestamp time.Time `json:"timestamp"`
}

func HealthHandler(startedAt time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{
			Status:        "ok",
			Service:       "api-herrmann13-portfolio",
			UptimeSeconds: int64(time.Since(startedAt).Seconds()),
			Timestamp:     time.Now().UTC(),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}
