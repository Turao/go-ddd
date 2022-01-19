package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/users/domain/user"
)

type UserRepository struct {
	users map[user.UserID]*user.User
}

var _ user.Repository = (*UserRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewUserRepository() (*UserRepository, error) {
	return &UserRepository{
		users: make(map[string]*user.User),
	}, nil
}

func (ur UserRepository) FindByID(ctx context.Context, id user.UserID) (*user.User, error) {
	t, found := ur.users[id]
	if !found {
		return nil, ErrNotFound
	}

	return t, nil
}

func (tr UserRepository) Save(ctx context.Context, p user.User) error {
	tr.users[p.ID] = &p
	return nil
}

// func (tr UserRepository) FindAll(ctx context.Context) ([]*user.User, error) {
// 	var ps []*user.User
// 	for _, p := range pr.users {
// 		ps = append(ps, p)
// 	}
// 	return ps, nil
// }
