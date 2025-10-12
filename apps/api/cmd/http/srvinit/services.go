package srvinit

import (
	"GoNext/base/internal/core/ports"
	"GoNext/base/internal/core/services/authservice"
	"GoNext/base/internal/core/services/exerciseservice"
	"GoNext/base/internal/core/services/userservice"
	"GoNext/base/internal/core/services/workoutexerciseservice"
	"GoNext/base/internal/core/services/workoutexercisesetservice"
	"GoNext/base/internal/core/services/workoutservice"
	"GoNext/base/pkg/config"
)

type Services struct {
	Auth               ports.AuthService
	User               ports.UserService
	Exercise           ports.ExerciseService
	Workout            ports.WorkoutService
	WorkoutExercise    ports.WorkoutExerciseService
	WorkoutExerciseSet    ports.WorkoutExerciseSetService
}

func InitServices(repos *Repositories, config *config.Config) *Services {
	return &Services{
		Auth:            authservice.NewAuthService(repos.User, config.Jwt.Secret),
		User:            userservice.NewUserService(repos.User),
		Workout:         workoutservice.NewWorkoutService(repos.Workout),
		Exercise:        exerciseservice.NewExerciseService(repos.Exercise),
		WorkoutExercise: workoutexerciseservice.NewWorkoutExerciseService(repos.WorkoutExercise),
		WorkoutExerciseSet: workoutexercisesetservice.NewWorkoutExerciseSetService(repos.WorkoutExerciseSet),
	}
}
