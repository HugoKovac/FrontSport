package srvinit

import (
	"GoNext/base/ent"
	"GoNext/base/internal/adapters/repositories"
	"GoNext/base/internal/core/ports"
)

type Repositories struct {
	User            ports.UserRepository
	Exercise        ports.ExerciseRepository
	Workout         ports.WorkoutRepository
	WorkoutExercise ports.WorkoutExerciseRepository
}

func InitRepos(entClient *ent.Client) *Repositories {
	return &Repositories{
		User:            repositories.NewUserRepository(entClient),
		Exercise:        repositories.NewExerciseRepository(entClient),
		Workout:         repositories.NewWorkoutRepository(entClient),
		WorkoutExercise: repositories.NewWorkoutExerciseRepository(entClient),
	}
}
