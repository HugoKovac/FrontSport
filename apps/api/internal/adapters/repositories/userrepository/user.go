package userrepository

import (
	"GoNext/base/ent"
	"GoNext/base/internal/core/ports"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) ports.UserRepository {
	return &UserRepository{
		client: client,
	}
}
