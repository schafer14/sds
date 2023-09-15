package examples

import (
	"context"
	"fmt"

	"github.com/schafer14/sds/mem"
	"github.com/segmentio/ksuid"
)

type user struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func (u user) GetID() string { return u.ID }

func Simple() int {

	// Setup the storage repository
	ctx := context.Background()
	userRepo := mem.New[user]()

	// Save an item
	id := ksuid.New().String()
	err := userRepo.Save(ctx, user{
		ID:    id,
		Name:  "Banner",
		Email: "banner@example.com",
	})
	if err != nil {
		return 1
	}

	// Retrieve the item
	user, err := userRepo.Find(ctx, id)
	if err != nil {
		return 1
	}

	// Check everything worked
	fmt.Println(user)
	return 0
}
