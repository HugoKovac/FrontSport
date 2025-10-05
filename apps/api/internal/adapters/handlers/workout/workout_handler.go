package workout

import (
	"GoNext/base/internal/core/ports"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R fiber.Router
	WorkoutService ports.WorkoutService
	Validate    *validator.Validate
}

type WorkoutHandler struct {
	WorkoutService ports.WorkoutService
	Validate    *validator.Validate
}

func New(c *Config) {
	h := &WorkoutHandler{
		WorkoutService: c.WorkoutService,
		Validate: c.Validate,
	}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *WorkoutHandler) {
	s.Get("/:workout_id", h.GetWorkout)
	s.Post("/", h.CreateWorkout)
}
