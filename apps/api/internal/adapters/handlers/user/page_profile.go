package user

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/views/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *UserHandler) UserProfilePage(c *fiber.Ctx) error {
	u := fibercontext.GetUserToContext(c)

	if u == nil || u.Id == uuid.Nil {
		c.Set("HX-Redirect", "/auth/register")
		return nil
	}

	use, err := h.UserService.GetById(u.Id)
	if err != nil {
		c.Set("HX-Redirect", "/auth/register")
		return nil
	}

	return templ.Render(c, user.Profile(use))
}
