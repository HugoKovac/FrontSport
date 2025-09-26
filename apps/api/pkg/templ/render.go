package templ

import (
	"context"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	cont := context.WithValue(c.Context(), "userID", c.Locals("userID"))
	return component.Render(cont, c.Response().BodyWriter())
}
