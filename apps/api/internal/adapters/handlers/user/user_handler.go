package user

import (
	"GoNext/base/internal/core/ports"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R           fiber.Router
	UserService ports.UserService
	Validate    *validator.Validate
}

type UserHandler struct {
	UserService ports.UserService
	Validate    *validator.Validate
}

func New(c *Config) {
	h := &UserHandler{
		UserService: c.UserService,
		Validate:    c.Validate,
	}

	registerUserRoute(c.R, h)
}

func registerUserRoute(s fiber.Router, h *UserHandler) {
	s.Get("/home", h.UserHomePage)
	s.Get("/profile", h.UserProfilePage)
	s.Put("/", h.UpdateUser)
}
