package dto

import "github.com/google/uuid"

type WorkoutExerciseSetQuery struct {
	WorkoutExerciseSetID uuid.UUID `params:"workout_exercise_set_id" validator:"required"`
}

type WorkoutExerciseSetUpdate struct {
	Weight int `json:"weight"`
	Reps   int `json:"reps"`
}
