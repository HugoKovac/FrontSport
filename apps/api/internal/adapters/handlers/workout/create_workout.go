package workout

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components"
	"GoNext/base/templ/views"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *WorkoutHandler) CreateWorkout(c *fiber.Ctx) error {
	u := c.Locals("userID")

	if u == nil {
		c.Set("HX-Redirect", "/auth/register")
		return nil
	}
	wrk, err := h.WorkoutService.CreateWorkout(uuid.MustParse(u.(string)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}
	c.Set("HX-Redirect", "/workout/" + wrk.Id)
	return templ.Render(c, views.Workout(time.Now().String()))
}
