package fibercontext

import (
	"GoNext/base/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

func SetUserToContext(c *fiber.Ctx, user *domain.User) {
	c.Locals("user", user)
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
