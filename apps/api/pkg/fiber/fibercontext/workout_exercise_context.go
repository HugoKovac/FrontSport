package fibercontext

import (
	"GoNext/base/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

func SetWorkoutExerciseToContext(c *fiber.Ctx, workoutExercise *domain.WorkoutExercise) {
	c.Locals("workout_exercise", workoutExercise)
}

func GetWorkoutExerciseToContext(c *fiber.Ctx) *domain.WorkoutExercise {
	workoutExercise, exists := c.Locals("workout_exercise").(*domain.WorkoutExercise)

	if !exists {
		c.Set("HX-Redirect", "/auth/register")
		c.Next()

		return nil
	}

	return workoutExercise
}
