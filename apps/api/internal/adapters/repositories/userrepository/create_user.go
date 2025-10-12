package userrepository

import (
	"GoNext/base/internal/core/domain"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

func (r *UserRepository) Create(user domain.User) (*domain.User, error) {
	ctx := context.Background()
	dUser, err := r.client.User.Create().
		SetEmail(user.Email).
		SetFirstname(user.Firstname).
		SetLastname(user.Lastname).
		SetPassword(user.Password).
		SetID(uuid.New()).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		log.Println("Creating User: ", err)
		return nil, err
	}

	return dUser.ToDomain(), nil
}
