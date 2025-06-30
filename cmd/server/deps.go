package main

import (
	"capsuler/internal/capsule"
	"capsuler/internal/sqlite"
	"database/sql"
	"net/http"
)

type dependencies struct {
	pool                 *sql.DB
	sqlCapsuleRepository capsule.Repository
	capsuleService       capsule.Service
	capsuleController    capsule.Controller
	capsuleHandler       http.Handler
}

func InitDependencies() *dependencies {
	pool := sqlite.NewConnection()
	sqlCapsuleRepository := capsule.NewRepository(pool)
	capsuleService := capsule.NewService(sqlCapsuleRepository)
	capsuleController := capsule.NewController(capsuleService)
	capsuleRoutes := capsule.NewHandler(capsuleController)

	return &dependencies{
		pool:                 pool,
		sqlCapsuleRepository: sqlCapsuleRepository,
		capsuleService:       capsuleService,
		capsuleController:    capsuleController,
		capsuleHandler:       capsuleRoutes,
	}
}
