package repository

import (
	"context"

	"github.com/o-t-k-t/ent-exercise/ent"
	"github.com/o-t-k-t/ent-exercise/ent/user"
	"github.com/o-t-k-t/ent-exercise/entity"
)

type UserRepository struct {
	entClient ent.Client
}

func (ur UserRepository) FindUser(id int) (entity.User, error) {
	u, err := ur.entClient.User.Get(context.Background(), id)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User(*u), nil
}

func (ur UserRepository) FindUsers() ([]entity.User, error) {
	us, err := ur.entClient.User.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	return ur.toEntitySlice(us), nil
}

func (ur UserRepository) FindUsersByName(name string) ([]entity.User, error) {
	us, err := ur.entClient.
		User.Query().Where(user.NameContains(name)).All(context.Background())
	if err != nil {
		return nil, err
	}

	return ur.toEntitySlice(us), nil
}

func (ur UserRepository) toEntitySlice(us []*ent.User) []entity.User {
	eus := make([]entity.User, len(us))
	for _, u := range us {
		eus = append(eus, entity.User(*u))
	}
	return eus
}
