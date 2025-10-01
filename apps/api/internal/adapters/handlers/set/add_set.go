package set

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components/set"

	"github.com/gofiber/fiber/v2"
)

func (h *SetHandler) AddSet(c *fiber.Ctx) error {
	return templ.Render(c, set.Card("Leg Press"))

}
