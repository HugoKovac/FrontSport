package error

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R fiber.Router
}

type ErrorHandler struct {
}

func New(c *Config) {
	h := &ErrorHandler{}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *ErrorHandler) {
	s.Get("/not-found", h.NotFoundPage)
}

