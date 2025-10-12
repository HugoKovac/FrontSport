package ent

import "GoNext/base/internal/core/domain"

func (r *WorkoutExerciseSet) ToDomain() (w *domain.WorkoutExerciseSet) {
	w = &domain.WorkoutExerciseSet{
		Id:         r.ID,
		Weight: r.Weight,
		Reps: r.Reps,
	}
	if r.Edges.WorkoutExercise != nil {
		w.WorkoutExercise =  r.Edges.WorkoutExercise.ToDomain()
	}
	return
}

func (r WorkoutExerciseSets) ToDomain() (exs []*domain.WorkoutExerciseSet) {
	for _, v := range r {
		exs = append(exs, v.ToDomain())
	}
	return
}
