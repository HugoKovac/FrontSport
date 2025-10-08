package fibercontext

import (
	"GoNext/base/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

func SetWorkoutToContext(c *fiber.Ctx, workout *domain.Workout) {
	c.Locals("workout", workout)
}

func GetWorkoutToContext(c *fiber.Ctx) *domain.Workout {
	workout, exists := c.Locals("workout").(*domain.Workout)

	if !exists {
		c.Set("HX-Redirect", "/auth/register")
		c.Next()

		return nil
	}

	return workout
}
