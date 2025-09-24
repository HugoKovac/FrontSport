package handlers

import (
	"GoNext/base/internal/core/ports"
	"GoNext/base/internal/core/services"
	"GoNext/base/internal/middleware"
	"GoNext/base/pkg/config"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app          *fiber.App
	authHandler  *AuthHandler
	userHandler  *UserHandler
	errorHandler *ErrorHandler
}

func NewRouter(app *fiber.App, userRepo ports.UserRepository, config *config.Config) *Router {

	authService := services.NewAuthService(userRepo, config.Jwt.Secret)
	userService := services.NewUserService(userRepo)

	authHandler := NewAuthHandler(authService, userService, config)
	userHandler := NewUserHandler(userService)
	errorHandler := NewErrorHandler()

	return &Router{app: app, authHandler: authHandler, userHandler: userHandler, errorHandler: errorHandler}
}

func (r *Router) SetupPublicRoutes() {
	// static
	r.app.Static("/assets", "./templ/assets")

	// rendered views
	r.app.Get("/auth/register", r.authHandler.RegisterPage)
	r.app.Get("/auth/login", r.authHandler.LoginPage)
	// r.app.Get("*", r.errorHandler.NotFoundPage)
	r.app.Get("/auth/register", r.authHandler.RegisterPage)


	// api
	r.app.Post("/api/auth/register", r.authHandler.Register)
	r.app.Post("/api/auth/login", r.authHandler.Login)
	r.app.Post("/api/auth/logout", r.authHandler.Logout)
	r.app.Get("/api/auth/status", r.authHandler.Status)
}

func (r *Router) SetupProtectedRoutes() {
	api := r.app.Group("/api", middleware.JWTAuthentication(r.authHandler.authService, true))
	api.Get("/users/me", r.userHandler.GetCurrentUser)
	api.Put("/users/me", r.userHandler.UpdateCurrentUser)
	api.Get("/users", r.userHandler.GetByEmail)
	
	// rendered views
	views := r.app.Group("/", middleware.JWTAuthentication(r.authHandler.authService, false))
	views.Get("/", r.userHandler.HomePage)
}
