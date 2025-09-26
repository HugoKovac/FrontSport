package auth

import (
	"GoNext/base/internal/adapters/dto"
	"GoNext/base/pkg/templ"
	customvalidator "GoNext/base/pkg/validator"
	"GoNext/base/templ/components"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var creds dto.UserCredentials
	if err := c.BodyParser(&creds); err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}

	if err := h.Validate.Struct(creds); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.Status(fiber.StatusUnprocessableEntity)
			return templ.Render(c, components.ErrorMessage(customvalidator.ErrorMessage(errs)))
		}
	}

	token, err := h.AuthService.Authenticate(creds.Email, creds.Password)
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		templ.Render(c, components.Input(components.InputAttributes{
			Id:          "password",
			Name:        "password",
			Type:        "password",
			Placeholder: "Password",
			Error:       true,
			OOB:         true,
		}))
		return templ.Render(c, components.ErrorMessage([]string{"Wrong credentials"}))
	}

	var domain string
	if os.Getenv("DEV") == "true" {
		domain = ".localhost"
	} else {
		domain = "." + h.Config.Env.Domain
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "Bearer " + token,
		Domain:   domain,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
	}
	c.Cookie(&cookie)
	c.Set("HX-Redirect", "/")
	return c.Status(fiber.StatusOK).Send([]byte{})
}
