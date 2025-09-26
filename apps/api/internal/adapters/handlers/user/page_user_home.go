package user

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/views"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) UserHomePage(c *fiber.Ctx) error {
	u := c.Locals("userID")

	if u == nil {
		c.Set("HX-Redirect", "/auth/register")
		return nil
	}

	user, err := h.UserService.GetById(u.(string))
	if err != nil {
		c.Set("HX-Redirect", "/auth/register")
		return nil
	}

	return templ.Render(c, views.ProtectedHome(user.Email))
}
