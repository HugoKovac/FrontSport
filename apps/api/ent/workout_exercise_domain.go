package ent

import "GoNext/base/internal/core/domain"

func (r *WorkoutExercise) ToDomain() (w *domain.WorkoutExercise) {
	w = &domain.WorkoutExercise{
		Id:         r.ID,
		ExerciseID: r.ExerciseID,
		WorkoutID:  r.WorkoutID,
	}
	if r.Edges.Exercise != nil {
		w.Exercise =  r.Edges.Exercise.ToDomain()
	}
	if r.Edges.Workout != nil {
		w.Workout = r.Edges.Workout.ToDomain()
	}
	return
}

func (r WorkoutExercises) ToDomain() (exs []*domain.WorkoutExercise) {
	for _, v := range r {
		exs = append(exs, v.ToDomain())
	}
	return
}
