package userrepository

import (
	"GoNext/base/ent/user"
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

func (r *UserRepository) FindById(id uuid.UUID) (*domain.User, error) {
	ctx := context.Background()
	u, err := r.client.User.Query().Where(user.ID(id)).Only(ctx)
	return u.ToDomain(), err
}
