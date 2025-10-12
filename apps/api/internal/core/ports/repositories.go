package ports

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user domain.User) (*domain.User, error)
	FindById(id uuid.UUID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
}

type ExerciseRepository interface {
	GetExercises() ([]*domain.Exercise, error)
	GetExerciseById(id int) (*domain.Exercise, error)
}

type WorkoutRepository interface {
	CreateWorkout(userId uuid.UUID) (*domain.Workout, error)
	GetWorkoutsByUser(userId uuid.UUID) ([]*domain.Workout, error)
	GetActiveWorkoutByUser(userId uuid.UUID) (*domain.Workout, error)
	GetWorkoutById(id uuid.UUID) (*domain.Workout, error)
	UpdateWorkoutToNotActive(id uuid.UUID) error
}

type WorkoutExerciseRepository interface {
	CreateWorkoutExercise(exerciseId int, userId uuid.UUID) (*domain.WorkoutExercise, error)
	GetWorkoutExercisesByWorkoutIdWithExAndSets(id uuid.UUID) ([]*domain.WorkoutExercise, error)
	GetWorkoutExerciseByIdFromUser(weID int, userId uuid.UUID) (*domain.WorkoutExercise, error)
}

type WorkoutExerciseSetRepository interface {
	CreateWorkoutExerciseSet(workoutExerciseId, weight, reps int) (*domain.WorkoutExerciseSet, error)
	GetWorkoutExerciseSetByWorkoutExerciseId(id int) ([]*domain.WorkoutExerciseSet, error)
	GetWorkoutExerciseSetByIdFromUser(id uuid.UUID, userID uuid.UUID) (*domain.WorkoutExerciseSet, error)
	UpdateWorkoutExerciseSet(id uuid.UUID, weight int, reps int) error
}
