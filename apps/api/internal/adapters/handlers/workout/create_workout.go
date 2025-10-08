package workout

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components"
	"GoNext/base/templ/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *WorkoutHandler) CreateWorkout(c *fiber.Ctx) error {
	u := fibercontext.GetUserToContext(c)
	wrk, err := h.WorkoutService.CreateWorkout(uuid.MustParse(u.Id))
	if err != nil {
		c.Status(422)
		return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Please finish your current workout to start a new one"}))
	}
	c.Set("HX-Redirect", "/workout/"+wrk.Id)
	return templ.Render(c, views.Workout(wrk.Id))
}
