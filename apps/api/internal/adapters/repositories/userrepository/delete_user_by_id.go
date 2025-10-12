package userrepository

import (
	"context"
	"log"

	"github.com/google/uuid"
)

func (r *UserRepository) Delete(id string) error {
	ctx := context.Background()
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return err
	}

	err = r.client.User.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		log.Println("Deleting User: ", err)
		return err
	}
	return nil
}
