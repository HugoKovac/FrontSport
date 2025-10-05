package workout

import (
	"GoNext/base/ent"
	"GoNext/base/internal/adapters/dto"
	"GoNext/base/pkg/templ"
	customvalidator "GoNext/base/pkg/validator"
	"GoNext/base/templ/components"
	"GoNext/base/templ/views"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *WorkoutHandler) GetWorkout(c *fiber.Ctx) error {
	var query dto.WorkoutQuery
	err := c.ParamsParser(&query)
	if err != nil {
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}
	if err := h.Validate.Struct(query); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.Status(fiber.StatusUnprocessableEntity)
			return templ.Render(c, components.ErrorMessage(customvalidator.ErrorMessage(errs)))
		}
	}
	if query.WorkoutID == uuid.Nil {
		c.Set("HX-Redirect", "/error/not-found")
		return nil
	}
	wrk, err := h.WorkoutService.GetWorkoutById(query.WorkoutID)
	if ent.IsNotFound(err) {
		c.Set("HX-Redirect", "/error/not-found")
		return nil
	}
	if err != nil {
		c.Status(500)
		return templ.Render(c, components.ErrorMessage([]string{"Internal Server Error"}))
	}
	return templ.Render(c, views.Workout(wrk.CreatedAt.String()))
}
