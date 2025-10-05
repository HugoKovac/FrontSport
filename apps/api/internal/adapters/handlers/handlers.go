package handlers

import (
	"GoNext/base/internal/adapters/handlers/auth"
	error_handler "GoNext/base/internal/adapters/handlers/error"
	"GoNext/base/internal/adapters/handlers/exercise"
	"GoNext/base/internal/adapters/handlers/public"
	"GoNext/base/internal/adapters/handlers/user"
	"GoNext/base/internal/adapters/handlers/workout"
	"GoNext/base/internal/core/ports"
	"GoNext/base/internal/core/services"
	"GoNext/base/internal/middleware"
	"GoNext/base/pkg/config"
	customvalidator "GoNext/base/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InitHandlers(app *fiber.App, userRepo ports.UserRepository, exerciseRepo ports.ExerciseRepository, workoutRepo ports.WorkoutRepository, config *config.Config) {

	authService := services.NewAuthService(userRepo, config.Jwt.Secret)
	userService := services.NewUserService(userRepo)
	workoutService := services.NewWorkoutService(workoutRepo)
	exerciseService := services.NewExerciseService(exerciseRepo)

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
		R:           global.Group("/"),
		WorkoutService: workoutService,
		UserService: userService,
	})
	workout.New(&workout.Config{
		R:              global.Group("/workout"),
		Validate:    v,
		WorkoutService: workoutService,
	})
	exercise.New(&exercise.Config{
		R:               global.Group("/exercise"),
		ExerciseService: exerciseService,
	})
	error_handler.New(&error_handler.Config{
		R: global.Group("/error"),
	})
}
