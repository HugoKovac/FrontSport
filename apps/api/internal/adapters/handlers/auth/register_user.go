package auth

import (
	"os"
	"time"

	"GoNext/base/internal/adapters/dto"
	"GoNext/base/internal/core/domain"
	"GoNext/base/pkg/templ"
	"GoNext/base/templ/components"

	customvalidator "GoNext/base/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var userDTO dto.RegisterCredentials
	if err := c.BodyParser(&userDTO); err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}

	if err := h.Validate.Struct(userDTO); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.Status(fiber.StatusUnprocessableEntity)
			return templ.Render(c, components.ErrorMessage(customvalidator.ErrorMessage(errs)))
		}
	}

	if userDTO.Password != userDTO.Confirm {
		c.Status(fiber.StatusUnprocessableEntity)
		templ.Render(c, components.Input(components.InputAttributes{
			Id:          "password",
			Name:        "password",
			Type:        "password",
			Placeholder: "Password",
			Error:       true,
			OOB:         true,
		}))
		templ.Render(c, components.Input(components.InputAttributes{
			Id:          "confirm",
			Name:        "confirm",
			Type:        "password",
			Placeholder: "Confirm Password",
			Error:       true,
			OOB:         true,
		}))
		return templ.Render(c, components.ErrorMessage([]string{"Password: password confirmation doesn't match"}))
	}

	_, err := h.UserService.Register(domain.User{
		Email:    userDTO.Email,
		Password: userDTO.Password,
	})
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}

	token, err := h.AuthService.Authenticate(userDTO.Email, userDTO.Password)
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return templ.Render(c, components.ErrorMessage([]string{"Invalid credentials"}))
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

	return c.Status(fiber.StatusCreated).Send([]byte{})
}
