package public

import (
	"GoNext/base/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R fiber.Router
	UserService ports.UserService
	WorkoutService ports.WorkoutService
}

type PublicHandler struct {
	UserService ports.UserService
	WorkoutService ports.WorkoutService
}

func New(c *Config) {
	h := &PublicHandler{
		UserService: c.UserService,
		WorkoutService: c.WorkoutService,
	}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *PublicHandler) {
	s.Get("/", h.HomePage)
	s.Get("/empty", h.Empty)
}
