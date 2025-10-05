package public

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *PublicHandler) HomePage(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID != nil {
		user, err := h.UserService.GetById(userID.(string))
		if err == nil {
			wrks, err := h.WorkoutService.GetWorkoutsByUser(uuid.MustParse(user.Id))
			if err == nil {
				return templ.Render(c, views.ProtectedHome(user.Firstname, wrks))
			}
		}
	}
	return templ.Render(c, views.PublicHome())

}
