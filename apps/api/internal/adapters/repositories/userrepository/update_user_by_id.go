package userrepository

import (
	"GoNext/base/internal/core/domain"
	"context"
	"time"
)

func (r *UserRepository) Update(user *domain.User) error {
	ctx := context.Background()

	return r.client.User.UpdateOneID(user.Id).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
}
