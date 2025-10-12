package workout

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components"

	"github.com/gofiber/fiber/v2"
)

func (h *WorkoutHandler) FinishWorkout(c *fiber.Ctx) (err error) {
	wrk := fibercontext.GetWorkoutToContext(c)
	wrk.WorkoutExercises, err = h.WorkoutExerciseService.GetWorkoutExercisesByWorkoutIdWithExAndSets(wrk.Id)
	if err != nil {
		return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Failed getting workout exercise"}))
	}

	if err := h.WorkoutService.UpdateWorkoutToNotActive(wrk.Id); err != nil {
		return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Failed finishing workout"}))
	}

	c.Set("HX-Redirect", "/")
	return nil
}
