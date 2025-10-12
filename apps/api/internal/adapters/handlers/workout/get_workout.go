package workout

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components"
	"GoNext/base/templ/components/exercise"
	"GoNext/base/templ/views"

	t "github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func (h *WorkoutHandler) GetWorkout(c *fiber.Ctx) (err error) {
	wrk := fibercontext.GetWorkoutToContext(c)
	wrk.WorkoutExercises, err = h.WorkoutExerciseService.GetWorkoutExercisesByWorkoutIdWithExAndSets(wrk.Id)
	if err != nil {
		return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Failed getting workout exercise"}))
	}
	var exercises []t.Component
	for _, we := range wrk.WorkoutExercises {
		if we.Exercise == nil {
			continue
		}
		var sets []t.Component
		for _, ws := range we.Sets {
			sets = append(sets,
				exercise.Set(exercise.SetAttributes{
					ID:             ws.Id.String(),
					PreviousWeight: 60,
					PreviousReps:   12,
					Weight:         ws.Weight,
					Reps:           ws.Reps,
				}),
			)
		}
		exercise := exercise.Card(we.Id, we.Exercise.Name, wrk.Active, sets...)

		exercises = append(exercises, exercise)
	}

	return templ.Render(c, views.Workout(wrk.Id, wrk.Active, exercises...))
}
