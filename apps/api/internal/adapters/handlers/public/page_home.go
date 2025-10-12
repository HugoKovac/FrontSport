package public

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/views"

	"github.com/gofiber/fiber/v2"
)

func (h *PublicHandler) HomePage(c *fiber.Ctx) error {
	user := fibercontext.GetUserToContext(c)
	if user != nil {
		user, err := h.UserService.GetById(user.Id)
		if err == nil {
			wrks, err := h.WorkoutService.GetWorkoutsByUser(user.Id)
			if err == nil {
				return templ.Render(c, views.ProtectedHome(user.Firstname, wrks))
			}
		}
	}
	return templ.Render(c, views.PublicHome())
}
