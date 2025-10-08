package ent

import "GoNext/base/internal/core/domain"

func (r *Workout) ToDomain() (w *domain.Workout) {
	w = &domain.Workout{
		Id:        r.ID.String(),
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Active:    r.Active,
	}
	if len(r.Edges.WorkoutExercise) > 0 {
		var we WorkoutExercises = r.Edges.WorkoutExercise
		w.WorkoutExercises = we.ToDomain()
	}
	return
}

func (r Workouts) ToDomain() (exs []*domain.Workout) {
	for _, v := range r {
		exs = append(exs, v.ToDomain())
	}
	return
}
