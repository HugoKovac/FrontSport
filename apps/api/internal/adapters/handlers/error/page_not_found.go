package error

import (
	"GoNext/base/pkg/templ"
	error_views "GoNext/base/templ/views/error"

	"github.com/gofiber/fiber/v2"
)

func (h *ErrorHandler) NotFoundPage(c *fiber.Ctx) error {
	return templ.Render(c, error_views.NotFound())
}
