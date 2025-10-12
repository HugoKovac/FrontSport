package srvinit

import (
	"GoNext/base/ent"
	"GoNext/base/internal/adapters/repositories/exerciserepository"
	"GoNext/base/internal/adapters/repositories/userrepository"
	"GoNext/base/internal/adapters/repositories/workoutexerciserepository"
	"GoNext/base/internal/adapters/repositories/workoutexercisesetrepository"
	"GoNext/base/internal/adapters/repositories/workoutrepository"
	"GoNext/base/internal/core/ports"
)

type Repositories struct {
	User               ports.UserRepository
	Exercise           ports.ExerciseRepository
	Workout            ports.WorkoutRepository
	WorkoutExercise    ports.WorkoutExerciseRepository
	WorkoutExerciseSet ports.WorkoutExerciseSetRepository
}

func InitRepos(entClient *ent.Client) *Repositories {
	return &Repositories{
		User:               userrepository.NewUserRepository(entClient),
		Exercise:           exerciserepository.NewExerciseRepository(entClient),
		Workout:            workoutrepository.NewWorkoutRepository(entClient),
		WorkoutExercise:    workoutexerciserepository.NewWorkoutExerciseRepository(entClient),
		WorkoutExerciseSet: workoutexercisesetrepository.NewWorkoutExerciseSetRepository(entClient),
	}
}
