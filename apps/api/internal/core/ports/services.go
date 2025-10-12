package ports

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

type UserService interface {
	Register(user domain.User) (*domain.User, error)
	GetById(id uuid.UUID) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(userID uuid.UUID, email string, oldPassword string, newPassword string) (domain.User, error)
	Delete(id string) error
}

type AuthService interface {
	Authenticate(username string, password string) (string, error)
	ValidateToken(tokenString string) (string, error)
}

type ExerciseService interface {
	GetExercises() ([]*domain.Exercise, error)
	GetExerciseById(id int) (*domain.Exercise, error)
}

type WorkoutService interface {
	CreateWorkout(userId uuid.UUID) (*domain.Workout, error)
	GetWorkoutsByUser(userId uuid.UUID) ([]*domain.Workout, error)
	GetActiveWorkoutByUser(userId uuid.UUID) (*domain.Workout, error)
	GetWorkoutById(id uuid.UUID) (*domain.Workout, error)
	UpdateWorkoutToNotActive(id uuid.UUID) error
}

type WorkoutExerciseService interface {
	CreateWorkoutExercise(exerciseId int, workoutId uuid.UUID) (*domain.WorkoutExercise, error)
	GetWorkoutExercisesByWorkoutIdWithExAndSets(id uuid.UUID) ([]*domain.WorkoutExercise, error)
	GetWorkoutExerciseByIdFromUser(weID int, userId uuid.UUID) (*domain.WorkoutExercise, error)
}

type WorkoutExerciseSetService interface {
	CreateWorkoutExerciseSet(workoutExerciseId, weight, reps int) (*domain.WorkoutExerciseSet, error)
	GetWorkoutExerciseSetByWorkoutExerciseId(id int) ([]*domain.WorkoutExerciseSet, error)
	GetWorkoutExerciseSetByIdFromUser(id uuid.UUID, userID uuid.UUID) (*domain.WorkoutExerciseSet, error)
	UpdateWorkoutExerciseSet(id uuid.UUID, weight int, reps int) error
}
