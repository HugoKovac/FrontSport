package auth

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	var domain string
	if os.Getenv("DEV") == "true" {
		domain = ".localhost"
	} else {
		domain = "." + h.Config.Env.Domain
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Domain:   domain,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
		Expires:  time.Now().Add(-time.Hour), // Past time to expire the cookie
		Path:     "/",
	}
	c.Cookie(&cookie)
	c.Set("HX-Refresh", "true")
	return nil
}
