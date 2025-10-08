package exercise

import (
	"GoNext/base/internal/adapters/handlers/middleware"
	"GoNext/base/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R                      fiber.Router
	ExerciseService        ports.ExerciseService
	WorkoutExerciseService ports.WorkoutExerciseService
	WorkoutService         ports.WorkoutService
}

type ExerciseHandler struct {
	ExerciseService        ports.ExerciseService
	WorkoutExerciseService ports.WorkoutExerciseService
	WorkoutService         ports.WorkoutService
}

func New(c *Config) {
	h := &ExerciseHandler{
		ExerciseService:        c.ExerciseService,
		WorkoutExerciseService: c.WorkoutExerciseService,
		WorkoutService:         c.WorkoutService,
	}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *ExerciseHandler) {
	s.Post("/add-set", h.AddSet)
	s.Post("/add/:workout_id",
		middleware.CheckWorkoutExists(h.WorkoutService),
		h.AddExercise)
	s.Get("/", h.GetExercises)
}
