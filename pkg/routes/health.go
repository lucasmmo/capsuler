package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func HealthCheck(pool *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		maxRetry := 3
		var err error
		for retry := range maxRetry {
			err = pool.Ping()
			if err == nil {
				break
			}
			retry++
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(map[string]any{
				"error": err.Error(),
			}); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
