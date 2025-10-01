package public

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/views"

	"github.com/gofiber/fiber/v2"
)

func (h *PublicHandler) HomePage(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID != nil {
		user, err := h.UserService.GetById(userID.(string))
		if err == nil {
			return templ.Render(c, views.ProtectedHome(user.Firstname))
		}
	}
	return templ.Render(c, views.PublicHome())

}
