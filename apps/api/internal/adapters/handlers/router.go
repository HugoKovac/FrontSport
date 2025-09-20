package handlers

import (
	"GoNext/base/internal/core/ports"
	"GoNext/base/internal/core/services"
	"GoNext/base/internal/middleware"
	"GoNext/base/pkg/config"
	"GoNext/base/templ/views/auth"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app         *fiber.App
	authHandler *AuthHandler
	userHandler *UserHandler
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func NewRouter(app *fiber.App, userRepo ports.UserRepository, config *config.Config) *Router {

	authService := services.NewAuthService(userRepo, config.Jwt.Secret)
	userService := services.NewUserService(userRepo)

	authHandler := NewAuthHandler(authService, userService, config)
	userHandler := NewUserHandler(userService)

	return &Router{app: app, authHandler: authHandler, userHandler: userHandler}
}

func (r *Router) SetupPublicRoutes() {
	// static
	r.app.Static("/templ", "./templ")

	// rendered views
	r.app.Get("/api/auth/register", func(ctx *fiber.Ctx) error {
		return Render(ctx, auth.Register())
	})

	// api
	r.app.Post("/api/auth/register", r.authHandler.Register)
	r.app.Post("/api/auth/login", r.authHandler.Login)
	r.app.Post("/api/auth/logout", r.authHandler.Logout)
	r.app.Get("/api/auth/status", r.authHandler.Status)
}

func (r *Router) SetupProtectedRoutes() {
	api := r.app.Group("/api", middleware.JWTAuthentication(r.authHandler.authService))
	api.Get("/users/me", r.userHandler.GetCurrentUser)
	api.Put("/users/me", r.userHandler.UpdateCurrentUser)
	api.Get("/users", r.userHandler.GetByEmail)
}
