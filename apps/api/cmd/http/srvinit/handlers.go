package srvinit

import (
	"GoNext/base/internal/adapters/handlers/auth"
	error_handler "GoNext/base/internal/adapters/handlers/error"
	"GoNext/base/internal/adapters/handlers/exercise"
	"GoNext/base/internal/adapters/handlers/public"
	"GoNext/base/internal/adapters/handlers/user"
	"GoNext/base/internal/adapters/handlers/workout"
	"GoNext/base/internal/adapters/handlers/middleware"
	"GoNext/base/pkg/config"
	customvalidator "GoNext/base/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InitHandlers(app *fiber.App, services *Services, config *config.Config) {

	v := validator.New()
	customvalidator.RegisterCustomValidators(v)

	app.Static("/assets", "./templ/assets")

	global := app.Group("/", middleware.JWTAuthentication(services.Auth))

	auth.New(&auth.Config{
		R:           global.Group("/auth"),
		AuthService: services.Auth,
		UserService: services.User,
		Validate:    v,
		Config:      config,
	})
	user.New(&user.Config{
		R:           global.Group("/users"),
		Validate:    v,
		UserService: services.User,
	})
	public.New(&public.Config{
		R:              global.Group("/"),
		UserService:    services.User,
		WorkoutService: services.Workout,
	})
	workout.New(&workout.Config{
		R:              global.Group("/workout"),
		Validate:       v,
		WorkoutService: services.Workout,
		WorkoutExerciseService: services.WorkoutExercise,
	})
	exercise.New(&exercise.Config{
		R:                      global.Group("/exercise"),
		ExerciseService:        services.Exercise,
		WorkoutService:         services.Workout,
		WorkoutExerciseService: services.WorkoutExercise,
		WorkoutExerciseSetService: services.WorkoutExerciseSet,
		Validate:    v,
	})
	error_handler.New(&error_handler.Config{
		R: global.Group("/error"),
	})
}
