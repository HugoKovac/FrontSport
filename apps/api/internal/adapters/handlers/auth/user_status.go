package auth

import "github.com/gofiber/fiber/v2"

func (h *AuthHandler) Status(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}

	claims, err := h.AuthService.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.JSON(fiber.Map{
		"user": claims,
	})
}
