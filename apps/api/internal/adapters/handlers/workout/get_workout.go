package workout

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components"
	"GoNext/base/templ/components/exercise"
	"GoNext/base/templ/views"

	t "github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *WorkoutHandler) GetWorkout(c *fiber.Ctx) (err error) {
	wrk := fibercontext.GetWorkoutToContext(c)
	wrk.WorkoutExercises, err = h.WorkoutExerciseService.GetWorkoutExerciseByWorkoutId(uuid.MustParse(wrk.Id))
	if err != nil {
		return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Failed getting workout exercise"}))
	}
	var exercises []t.Component
	for _, v := range wrk.WorkoutExercises {
		if v.Exercise == nil {
			continue
		}
		exercise := exercise.Card(-1, v.Exercise.Name, exercise.Set(exercise.SetAttributes{
			Index:          1,
			PreviousWeight: 60,
			PreviousReps:   12,
		}))

		exercises = append(exercises, exercise)
	}

	return templ.Render(c, views.Workout(wrk.Id, exercises...))
}
