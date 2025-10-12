package dto

type WorkoutExerciseQuery struct {
	WorkoutExerciseID int `params:"workout_exercise_id" validator:"required"`
}
