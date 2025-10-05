package exercise

import (
	"GoNext/base/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R fiber.Router
	ExerciseService ports.ExerciseService
}

type ExerciseHandler struct {
	ExerciseService ports.ExerciseService
}

func New(c *Config) {
	h := &ExerciseHandler{
		ExerciseService: c.ExerciseService,
	}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *ExerciseHandler) {
	s.Post("/add-set", h.AddSet)
	s.Post("/add", h.AddExercise)
	s.Get("/", h.GetExercises)
}
