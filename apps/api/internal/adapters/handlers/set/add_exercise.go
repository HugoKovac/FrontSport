package set

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components/set"

	"github.com/gofiber/fiber/v2"
)

func (h *SetHandler) AddExercise(c *fiber.Ctx) error {
	return templ.Render(c, set.Exercise(set.ExerciseAttributes{
		Index: 4,
		PreviousWeight: 60,
		PreviousReps: 12,
	}))

}
