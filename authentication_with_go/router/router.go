package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router, RoleRouter Router) *chi.Mux {
	
	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.Logger)

	// chiRouter.Use(middlewares.RequestLogger) // Self created middleware for logging requests
	chiRouter.Use(middlewares.RateLimiterMiddleware)

	chiRouter.Get("/ping", controllers.PingHandler)

	chiRouter.HandleFunc("/fakestore/*", utils.ProxyToService("https://fakestoreapi.com/", "/fakestore"))

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)
	
	return chiRouter
}