package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/iagomaia/dload-tech-challenge/internal/infra/server/middlewares"
	"github.com/iagomaia/dload-tech-challenge/internal/infra/server/routes"
	"github.com/iagomaia/dload-tech-challenge/internal/infra/utils"
)

func GetServerRoutes() *chi.Mux {
	r := chi.NewRouter()
	logger := utils.GetLogger()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middlewares.CidMiddleware)
	r.Use(middlewares.GetCorsMiddleware())
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middlewares.ApplicationHeaders)
	r.Mount("/user/public", routes.GetPublicUserRoutes())
	r.Mount("/user", routes.GetUserRoutes())
	return r
}
