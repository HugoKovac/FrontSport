package exercise

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components/exercise"

	"github.com/gofiber/fiber/v2"
)

func (h *ExerciseHandler) AddExercise(c *fiber.Ctx) error {
	return templ.Render(c, exercise.Set(exercise.SetAttributes{
		Index: 4,
		PreviousWeight: 60,
		PreviousReps: 12,
	}))

}
