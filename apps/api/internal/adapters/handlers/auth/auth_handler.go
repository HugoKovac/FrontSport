package auth

import (
	"GoNext/base/internal/core/ports"
	"GoNext/base/pkg/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R           fiber.Router
	AuthService ports.AuthService
	UserService ports.UserService
	Validate    *validator.Validate
	Config      *config.Config
}

type AuthHandler struct {
	AuthService ports.AuthService
	UserService ports.UserService
	Validate    *validator.Validate
	Config      *config.Config
}

func New(c *Config) {
	h := &AuthHandler{
		AuthService: c.AuthService,
		UserService: c.UserService,
		Validate:    c.Validate,
		Config:      c.Config,
	}
	registerAuthRoute(c.R, h)
}

func registerAuthRoute(s fiber.Router, h *AuthHandler) {
	s.Get("/login", h.LoginPage)
	s.Get("/register", h.RegisterPage)

	s.Get("/status", h.Status)

	s.Post("/register", h.Register)
	s.Post("/login", h.Login)
	s.Post("/logout", h.Logout)
}
