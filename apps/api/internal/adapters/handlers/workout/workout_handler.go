package workout

import (
	"GoNext/base/internal/adapters/handlers/middleware"
	"GoNext/base/internal/core/ports"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R                      fiber.Router
	WorkoutService         ports.WorkoutService
	WorkoutExerciseService ports.WorkoutExerciseService
	Validate               *validator.Validate
}

type WorkoutHandler struct {
	WorkoutService         ports.WorkoutService
	WorkoutExerciseService ports.WorkoutExerciseService
	Validate               *validator.Validate
}

func New(c *Config) {
	h := &WorkoutHandler{
		WorkoutService:         c.WorkoutService,
		WorkoutExerciseService: c.WorkoutExerciseService,
		Validate:               c.Validate,
	}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *WorkoutHandler) {
	s.Get("/:workout_id",
		middleware.CheckWorkoutExists(h.WorkoutService),
		h.GetWorkout)
	s.Post("/", h.CreateWorkout)
	s.Post("/:workout_id/finish",
		middleware.CheckWorkoutExists(h.WorkoutService),
		h.FinishWorkout)
}
