package handlers

import (
	"GoNext/base/internal/adapters/handlers/auth"
	error_handler "GoNext/base/internal/adapters/handlers/error"
	"GoNext/base/internal/adapters/handlers/public"
	"GoNext/base/internal/adapters/handlers/user"
	"GoNext/base/internal/core/ports"
	"GoNext/base/internal/core/services"
	"GoNext/base/internal/middleware"
	"GoNext/base/pkg/config"
	customvalidator "GoNext/base/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InitHandlers(app *fiber.App, userRepo ports.UserRepository, config *config.Config) {

	authService := services.NewAuthService(userRepo, config.Jwt.Secret)
	userService := services.NewUserService(userRepo)

	v := validator.New()
	customvalidator.RegisterCustomValidators(v)

	app.Static("/assets", "./templ/assets")

	global := app.Group("/", middleware.JWTAuthentication(authService))

	auth.New(&auth.Config{
		R:           global.Group("/auth"),
		AuthService: authService,
		UserService: userService,
		Validate:    v,
		Config:      config,
	})
	user.New(&user.Config{
		R:           global.Group("/users"),
		Validate:    v,
		UserService: userService,
	})
	public.New(&public.Config{
		R: global.Group("/"),
	})
	error_handler.New(&error_handler.Config{
		R: global.Group("/error"),
	})
}
