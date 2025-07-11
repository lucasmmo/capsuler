package main

import (
	capsule "capsuler/internal/domain/capsule/model"
	payment "capsuler/internal/domain/payment/model"
	user "capsuler/internal/domain/user/model"
	userServices "capsuler/internal/domain/user/services"
	"capsuler/internal/infra/controllers"
	"capsuler/internal/infra/datasources"
	"capsuler/internal/infra/routes"
	"capsuler/web/templates"
	"context"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		fx.Provide(NewGormSQLiteConnectionDB),
		fx.Provide(datasources.NewSqlUserRepository),
		fx.Provide(templates.NewTemplate),
		fx.Provide(userServices.NewLoginService),
		fx.Provide(controllers.NewLoginController),
		fx.Provide(userServices.NewRegisterService),
		fx.Provide(controllers.NewRegisterController),
		fx.Provide(controllers.NewLandingPageController),
		fx.Provide(controllers.NewCapsuleDashboardController),
		fx.Provide(routes.NewRoutes),
		fx.Invoke(NewHTTPServer),
	).Run()
}

func NewGormSQLiteConnectionDB(lc fx.Lifecycle) *gorm.DB {
	dsn := "capsuler.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.AutoMigrate(&user.User{}, &capsule.Capsule{}, &capsule.Message{}, &payment.Payment{}); err != nil {
				return err
			}
			return nil
		},
	})
	return db
}

func NewGormPostgreSQLConnectionDB(lc fx.Lifecycle) *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.AutoMigrate(&user.User{}, &capsule.Capsule{}, &payment.Payment{}); err != nil {
				return err
			}
			return nil
		},
	})
	return db
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting HTTP server at :8080")
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Close()
		},
	})
	return srv
}
