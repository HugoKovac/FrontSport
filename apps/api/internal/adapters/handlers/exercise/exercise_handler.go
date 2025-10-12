package exercise

import (
	"GoNext/base/internal/adapters/handlers/middleware"
	"GoNext/base/internal/core/ports"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R                         fiber.Router
	ExerciseService           ports.ExerciseService
	WorkoutExerciseService    ports.WorkoutExerciseService
	WorkoutExerciseSetService ports.WorkoutExerciseSetService
	WorkoutService            ports.WorkoutService
	Validate                  *validator.Validate
}

type ExerciseHandler struct {
	ExerciseService           ports.ExerciseService
	WorkoutExerciseService    ports.WorkoutExerciseService
	WorkoutExerciseSetService ports.WorkoutExerciseSetService
	WorkoutService            ports.WorkoutService
	Validate                  *validator.Validate
}

func New(c *Config) {
	h := &ExerciseHandler{
		ExerciseService:           c.ExerciseService,
		WorkoutExerciseService:    c.WorkoutExerciseService,
		WorkoutService:            c.WorkoutService,
		WorkoutExerciseSetService: c.WorkoutExerciseSetService,
		Validate:                  c.Validate,
	}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *ExerciseHandler) {
	s.Post("/set/:workout_exercise_id",
		middleware.CheckWorkoutExerciseExists(h.WorkoutExerciseService),
		h.AddSet)
	s.Post("/add/:workout_id",
		middleware.CheckWorkoutExists(h.WorkoutService),
		h.AddExercise)
	s.Get("/", h.GetExercises)
	s.Patch("/set/:workout_exercise_set_id",
		middleware.CheckWorkoutExerciseSetExists(h.WorkoutExerciseSetService),
		h.PatchUpdateSet)
}
