package public

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/views"

	"github.com/gofiber/fiber/v2"
)

func (h *PublicHandler) Empty(c *fiber.Ctx) error {
	return templ.Render(c, views.Empty())
}
