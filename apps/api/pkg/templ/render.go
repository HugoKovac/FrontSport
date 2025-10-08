package templ

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"context"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	cont := context.WithValue(c.Context(), "user", fibercontext.GetUserToContext(c))
	return component.Render(cont, c.Response().BodyWriter())
}
