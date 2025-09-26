package public

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	R fiber.Router
}

type PublicHandler struct {
}

func New(c *Config) {
	h := &PublicHandler{}

	registerPublicRoute(c.R, h)
}

func registerPublicRoute(s fiber.Router, h *PublicHandler) {
	s.Get("/", h.HomePage)
}
