package auth

import (
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/views/auth"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) RegisterPage(c *fiber.Ctx) error {
	return templ.Render(c, auth.Register())
}