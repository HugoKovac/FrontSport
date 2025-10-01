package exercise

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components/exercise/modal"

	"github.com/gofiber/fiber/v2"
)

func (h *ExerciseHandler) GetExercises(c *fiber.Ctx) error {
	exs, err := h.ExerciseService.GetExercises()
	if err != nil {
		return err
	}
	return templ.Render(c, modal.List(exs))
}
