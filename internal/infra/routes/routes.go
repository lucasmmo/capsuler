package routes

import (
	"capsuler/internal/infra/controllers"
	"capsuler/internal/infra/middlewares"
	"net/http"

	"go.uber.org/fx"
)

func NewRoutes(
	lc fx.Lifecycle,
	loginController *controllers.LoginUser,
	registerController *controllers.RegisterUser,
	landingPageController *controllers.LandingPage,
	capsuleDashboardController *controllers.CapsuleDashboard,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", landingPageController.LandingPage)
	mux.HandleFunc("/login", loginController.Login)
	mux.HandleFunc("/register", registerController.Register)
	mux.HandleFunc("/capsules/dashboard", middlewares.AuthMiddleware(capsuleDashboardController.Dashboard))
	return mux
}
