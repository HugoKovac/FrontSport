package exercise

import (
	"GoNext/base/internal/adapters/dto"
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components"
	exercise_comp "GoNext/base/templ/components/exercise"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *ExerciseHandler) AddExercise(c *fiber.Ctx) error {
	query := &dto.ExerciseQuery{}
	wrk := fibercontext.GetWorkoutToContext(c)
	err := c.BodyParser(query)
	if err != nil {
		return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: err.Error()}))
	}
	for _, v := range query.Exercise {
		exercise, err := h.ExerciseService.GetExerciseById(v)
		if err != nil {
			log.Println(err.Error())
			return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Failed creating exercise"}))
		}
		if _, err = h.WorkoutExerciseService.CreateWorkoutExercise(exercise.Id, uuid.MustParse(wrk.Id)); err != nil {
			log.Println(err.Error())
			return templ.Render(c, components.Toast(components.ToastAttributes{T: "error", Message: "Failed creating workout exercise"}))
		}
		templ.Render(c, exercise_comp.Card(exercise.Id, exercise.Name, exercise_comp.Set(exercise_comp.SetAttributes{
			Index:          1,
			PreviousWeight: 60,
			PreviousReps:   12,
		})))
	}
	return nil
}
