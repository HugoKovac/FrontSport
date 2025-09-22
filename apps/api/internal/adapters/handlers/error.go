package handlers

import (
	"GoNext/base/pkg/templ"
	error_views "GoNext/base/templ/views/error"

	"github.com/gofiber/fiber/v2"
)

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h *ErrorHandler) NotFoundPage(c *fiber.Ctx) error {
	return templ.Render(c, error_views.NotFound())
}
