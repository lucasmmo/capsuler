package main

import (
	"capsuler/pkg/repositories"
	"capsuler/pkg/routes"
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	pool, err := sql.Open("sqlite3", "./capsules.db")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	if _, err := pool.Exec(`
		CREATE TABLE IF NOT EXISTS capsules (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			date_to_open DATETIME NOT NULL,
			is_open BOOLEAN NOT NULL DEFAULT FALSE,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		)
	`); err != nil {
		panic(err)
	}

	if _, err := pool.Exec(`
		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			capsule_id TEXT NOT NULL,
			message TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			FOREIGN KEY (capsule_id) REFERENCES capsules(id)
		)
	`); err != nil {
		panic(err)
	}

	sqlCapsuleRepository := repositories.NewSqlCapsuleRepository(pool)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", routes.HealthCheck(pool))
	mux.HandleFunc("POST /capsules", routes.CreateCapsule(sqlCapsuleRepository))
	mux.HandleFunc("POST /capsules/{id}/open", routes.OpenCapsule(sqlCapsuleRepository))
	mux.HandleFunc("POST /capsules/{id}/messages", routes.AddMessageToCapsule(sqlCapsuleRepository))

	http.ListenAndServe(":8080", mux)
}
