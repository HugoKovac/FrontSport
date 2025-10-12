package exercise

import (
	"GoNext/base/internal/adapters/dto"
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	customvalidator "GoNext/base/pkg/validator"
	"GoNext/base/templ/components"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h *ExerciseHandler) PatchUpdateSet(c *fiber.Ctx) error {
	ws := fibercontext.GetWorkoutExerciseSetToContext(c)
	var query dto.WorkoutExerciseSetUpdate
	err := c.BodyParser(&query)
	if err != nil {
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}
	if err := h.Validate.Struct(query); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.Status(fiber.StatusUnprocessableEntity)
			return templ.Render(c, components.ErrorMessage(customvalidator.ErrorMessage(errs)))
		}
	}
	if err = h.WorkoutExerciseSetService.UpdateWorkoutExerciseSet(ws.Id, query.Weight, query.Reps); err != nil {
		log.Println(err.Error())
		return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Failed updating set"}))
	}

	c.Status(201)
	return nil
}
