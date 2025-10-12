package userrepository

import (
	"GoNext/base/ent/user"
	"GoNext/base/internal/core/domain"
	"context"
)

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	ctx := context.Background()
	u, err := r.client.User.Query().Where(user.Email(email)).Only(ctx)
	return u.ToDomain(), err
}
