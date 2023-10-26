package routes

import (
	"github.com/go-chi/chi/v5"
	usercontracts "github.com/iagomaia/dload-tech-challenge/internal/domain/contracts/user"
	usercontrollersfactories "github.com/iagomaia/dload-tech-challenge/internal/factories/controllers/user"
	"github.com/iagomaia/dload-tech-challenge/internal/infra/adapters"
	"github.com/iagomaia/dload-tech-challenge/internal/infra/server/middlewares"
)

func GetPublicUserRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", adapters.AdaptRoute[usercontracts.SignupRequest](usercontrollersfactories.GetUserSignupController(), nil))
	r.Post("/login", adapters.AdaptRoute[usercontracts.LoginRequest](usercontrollersfactories.GetUserLoginController(), nil))
	return r
}

func GetUserRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.AuthMiddleware)
	r.Get("/me", adapters.AdaptRoute[any](usercontrollersfactories.GetMeController(), nil))
	r.Delete("/me", adapters.AdaptRoute[any](usercontrollersfactories.GetDeleteMeController(), nil))
	r.Patch("/{userId}", adapters.AdaptRoute[usercontracts.UpdateUserRequest](usercontrollersfactories.GetUpdateUserController(), &[]string{"userId"}))
	return r
}
