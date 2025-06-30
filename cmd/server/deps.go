package main

import (
	"capsuler/internal/capsule"
	"capsuler/internal/sqlite"
	"database/sql"
	"net/http"
)

type dependencies struct {
	pool          *sql.DB
	capsuleRoutes http.Handler
}

func InitDependencies() *dependencies {
	pool := sqlite.NewConnection()
	sqlCapsuleRepository := capsule.NewRepository(pool)

	capsuleCreator := capsule.NewCreator(sqlCapsuleRepository)
	capsuleOpener := capsule.NewOpener(sqlCapsuleRepository)
	capsuleMessageAdder := capsule.NewMessageAdder(sqlCapsuleRepository)
	capsuleController := capsule.NewController(capsuleCreator, capsuleOpener, capsuleMessageAdder)

	capsuleRoutes := capsule.NewRouter(capsuleController)

	return &dependencies{
		pool:          pool,
		capsuleRoutes: capsuleRoutes,
	}
}
