package srvinit

import (
	"GoNext/base/internal/core/ports"
	"GoNext/base/internal/core/services"
	"GoNext/base/pkg/config"
)

type Services struct {
	Auth            ports.AuthService
	User            ports.UserService
	Exercise        ports.ExerciseService
	Workout         ports.WorkoutService
	WorkoutExercise ports.WorkoutExerciseService
}


func InitServices(repos *Repositories, config *config.Config) *Services {
	return &Services{
		Auth:            services.NewAuthService(repos.User, config.Jwt.Secret),
		User:            services.NewUserService(repos.User),
		Workout:         services.NewWorkoutService(repos.Workout),
		Exercise:        services.NewExerciseService(repos.Exercise),
		WorkoutExercise: services.NewWorkoutExerciseService(repos.WorkoutExercise),
	}
}
