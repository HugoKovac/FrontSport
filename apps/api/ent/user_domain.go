package ent

import "GoNext/base/internal/core/domain"

func (r *User) ToDomain() *domain.User {
	if r == nil {
		return nil
	}

	return &domain.User{
		Id:        r.ID.String(),
		Firstname: r.Firstname,
		Lastname:  r.Lastname,
		Email:     r.Email,
		Password:  r.Password,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.CreatedAt,
		Role:      r.Role.String(),
	}
}
