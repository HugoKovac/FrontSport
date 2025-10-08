package middleware

import (
	"GoNext/base/ent"
	"GoNext/base/internal/adapters/dto"
	"GoNext/base/internal/core/ports"
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	customvalidator "GoNext/base/pkg/validator"
	"GoNext/base/templ/components"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CheckWorkoutExists(workoutService ports.WorkoutService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query dto.WorkoutQuery
		v := validator.New()
		err := c.ParamsParser(&query)
		if err != nil {
			return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
		}
		if err := v.Struct(query); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				c.Status(fiber.StatusUnprocessableEntity)
				return templ.Render(c, components.ErrorMessage(customvalidator.ErrorMessage(errs)))
			}
		}
		if query.WorkoutID == uuid.Nil {
			c.Set("HX-Redirect", "/error/not-found")
			return nil
		}
		wrk, err := workoutService.GetWorkoutById(query.WorkoutID)
		if ent.IsNotFound(err) {
			c.Set("HX-Redirect", "/error/not-found")
			return nil
		}
		fibercontext.SetWorkoutToContext(c, wrk)
		return c.Next()
	}
}
