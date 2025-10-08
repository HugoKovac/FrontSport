package middleware

import (
	"strings"

	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"GoNext/base/pkg/fiber/fibercontext"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(authService ports.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		if token != "" {
			parts := strings.Split(token, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				userID, err := authService.ValidateToken(token)
				if err == nil {
					// Set user ID in context for use in protected routes
					fibercontext.SetUserToContext(c, &domain.User{
						Id: userID,
					})
					return c.Next()
				}
			}
		}

		if c.Method() == "GET" {
			c.Set("HX-Redirect", "/auth/register")
			return c.Next()
		} else {
			fibercontext.SetUserToContext(c, &domain.User{})
			return c.Next()
		}
	}
}
