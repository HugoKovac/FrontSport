package set

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R fiber.Router
}

type SetHandler struct {
}

func New(c *Config) {
	h := &SetHandler{}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *SetHandler) {
	s.Post("/add-exercise", h.AddExercise)
	s.Post("/add", h.AddSet)
}
