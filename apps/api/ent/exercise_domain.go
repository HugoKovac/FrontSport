package ent

import "GoNext/base/internal/core/domain"

func (r *Exercise) ToDomain() *domain.Exercise {
	return &domain.Exercise{
		Id:        r.ID,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Name:      r.Name,
	}
}

func (r Exercises) ToDomain() (exs []*domain.Exercise) {
	for _, v := range r {
		exs = append(exs, v.ToDomain())
	}
	return
}
