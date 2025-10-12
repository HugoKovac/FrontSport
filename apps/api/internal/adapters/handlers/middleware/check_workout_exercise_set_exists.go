package middleware

import (
	"GoNext/base/ent"
	"GoNext/base/internal/adapters/dto"
	"GoNext/base/internal/core/ports"
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	customvalidator "GoNext/base/pkg/validator"
	"GoNext/base/templ/components"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CheckWorkoutExerciseSetExists(workoutExerciseSetService ports.WorkoutExerciseSetService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query dto.WorkoutExerciseSetQuery
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
		user := fibercontext.GetUserToContext(c)
		wes, err := workoutExerciseSetService.GetWorkoutExerciseSetByIdFromUser(query.WorkoutExerciseSetID, user.Id)
		if ent.IsNotFound(err) {
			log.Println(err)
			c.Set("HX-Redirect", "/error/not-found")
			return nil
		}
		if err != nil {
			log.Println(err.Error())
			return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Fail getting this set"}))
		}
		fibercontext.SetWorkoutExerciseSetToContext(c, wes)
		return c.Next()
	}
}
