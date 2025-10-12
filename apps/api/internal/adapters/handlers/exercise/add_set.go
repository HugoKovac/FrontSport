package exercise

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components"
	"GoNext/base/templ/components/exercise"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *ExerciseHandler) AddSet(c *fiber.Ctx) error {
	we := fibercontext.GetWorkoutExerciseToContext(c)
	wes, err := h.WorkoutExerciseSetService.CreateWorkoutExerciseSet(we.Id, 0, 0) // update when check
	if err != nil {
		log.Println(err.Error())
		return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Failed creating set"}))
	}
	return templ.Render(c, exercise.Set(exercise.SetAttributes{
		WorkoutExerciseId: we.Id,
		ID:          wes.Id.String(),
		PreviousWeight: 0,
		PreviousReps:   0,
	}))

}
