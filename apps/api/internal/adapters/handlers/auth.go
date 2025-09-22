// internal/adapters/handlers/auth_handler.go
package handlers

import (
	"GoNext/base/internal/adapters/dto"
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"GoNext/base/pkg/config"
	"GoNext/base/pkg/templ"
	customvalidator "GoNext/base/pkg/validator"
	"GoNext/base/templ/components"
	"GoNext/base/templ/views/auth"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService ports.AuthService
	userService ports.UserService
	validate    *validator.Validate
	config      *config.Config
}

func NewAuthHandler(authService ports.AuthService, userService ports.UserService, config *config.Config) *AuthHandler {
	v := validator.New()
	customvalidator.RegisterCustomValidators(v)

	return &AuthHandler{
		authService: authService,
		userService: userService,
		validate:    v,
		config:      config,
	}
}

func (h *AuthHandler) RegisterPage(c *fiber.Ctx) error {
	return templ.Render(c, auth.Register())
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var userDTO dto.UserCredentials
	if err := c.BodyParser(&userDTO); err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}

	log.Println(userDTO)

	if err := h.validate.Struct(userDTO); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.Status(fiber.StatusUnprocessableEntity)
			return templ.Render(c, components.ErrorMessage(customvalidator.ErrorMessage(errs)))
		}
	}

	if userDTO.Password != userDTO.Confirm {
		c.Status(fiber.StatusUnprocessableEntity)
		templ.Render(c, components.Input(components.InputAttributes{
                                Id: "password",
                                Name: "password",
                                Type: "password",
								Placeholder: "Password",
								Error: true,
								OOB: true,
                            }))
		templ.Render(c, components.Input(components.InputAttributes{
                                Id: "confirm",
                                Name: "confirm",
                                Type: "password",
								Placeholder: "Confirm Password",
								Error: true,
								OOB: true,
                            }))
		return templ.Render(c, components.ErrorMessage([]string{"Password: password confirmation doesn't match"}))
	}

	_, err := h.userService.Register(domain.User{
		Email:    userDTO.Email,
		Password: userDTO.Password,
	})
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}

	token, err := h.authService.Authenticate(userDTO.Email, userDTO.Password)
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return templ.Render(c, components.ErrorMessage([]string{"Invalid credentials"}))
	}

	var domain string
	if os.Getenv("DEV") == "true" {
		domain = ".localhost"
	} else {
		domain = "." + h.config.Env.Domain
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
	c.Set("HX-Redirect", "/api/protected-home")

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var creds dto.UserCredentials
	if err := c.BodyParser(&creds); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if err := h.validate.Struct(creds); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	token, err := h.authService.Authenticate(creds.Email, creds.Password)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	var domain string
	if os.Getenv("DEV") == "true" {
		domain = ".localhost"
	} else {
		domain = "." + h.config.Env.Domain
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

func (h *AuthHandler) Status(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}

	claims, err := h.authService.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.JSON(fiber.Map{
		"user": claims,
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	var domain string
	if os.Getenv("DEV") == "true" {
		domain = ".localhost"
	} else {
		domain = "." + h.config.Env.Domain
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
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
