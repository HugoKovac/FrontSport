package fibercontext

import (
	"GoNext/base/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

func SetUserToContext(c *fiber.Ctx, workout *domain.User) {
	c.Locals("user", workout)
}

func GetUserToContext(c *fiber.Ctx) *domain.User {
	user, exists := c.Locals("user").(*domain.User)

	if !exists {
		c.Set("HX-Redirect", "/auth/register")
		c.Next()
		return nil
	}

	return user
}
