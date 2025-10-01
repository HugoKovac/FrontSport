package repositories

import (
	"GoNext/base/ent"
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"context"
)

type ExerciseRepository struct {
	client *ent.Client
}

func NewExerciseRepository(client *ent.Client) ports.ExerciseRepository {
	return &ExerciseRepository{
		client: client,
	}
}

func (r *ExerciseRepository) ToDomainExercise(entExercise *ent.Exercise) *domain.Exercise {
	return &domain.Exercise{
		Id:        entExercise.ID,
		CreatedAt: entExercise.CreatedAt,
		UpdatedAt: entExercise.UpdatedAt,
		Name:      entExercise.Name,
	}
}

func (r *ExerciseRepository) ToDomainExercises(entExercises []*ent.Exercise) (exs []*domain.Exercise) {
	for _, v := range entExercises {
		exs = append(exs, r.ToDomainExercise(v))
	}
	return
}

func (r *ExerciseRepository) GetExercises() ([]*domain.Exercise, error) {
	ctx := context.Background()
	exercises, err := r.client.Exercise.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return r.ToDomainExercises(exercises), nil

}
