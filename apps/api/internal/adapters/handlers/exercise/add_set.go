package exercise

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components/exercise"

	"github.com/gofiber/fiber/v2"
)

func (h *ExerciseHandler) AddSet(c *fiber.Ctx) error {
	return templ.Render(c, exercise.Card("Leg Press"))

}
