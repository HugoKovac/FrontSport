package user

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/views/user"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) UserProfilePage(c *fiber.Ctx) error {
	u := c.Locals("userID")

	if u == nil {
		c.Set("HX-Redirect", "/auth/register")
		return nil
	}

	use, err := h.UserService.GetById(u.(string))
	if err != nil {
		c.Set("HX-Redirect", "/auth/register")
		return nil
	}

	return templ.Render(c, user.Profile(use))
}
