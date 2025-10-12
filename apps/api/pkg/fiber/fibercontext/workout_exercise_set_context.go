package fibercontext

import (
	"GoNext/base/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

func SetWorkoutExerciseSetToContext(c *fiber.Ctx, workoutExercise *domain.WorkoutExerciseSet) {
	c.Locals("workout_exercise_set", workoutExercise)
}

func GetWorkoutExerciseSetToContext(c *fiber.Ctx) *domain.WorkoutExerciseSet {
	workoutExercise, exists := c.Locals("workout_exercise_set").(*domain.WorkoutExerciseSet)

	if !exists {
		c.Set("HX-Redirect", "/auth/register")
		c.Next()

		return nil
	}

	return workoutExercise
}
