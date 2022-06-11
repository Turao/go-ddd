package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/users/domain/user"
)

type UserRepository struct {
	users map[user.UserID]*user.UserAggregate
}

var _ user.Repository = (*UserRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewUserRepository() (*UserRepository, error) {
	return &UserRepository{
		users: make(map[string]*user.UserAggregate),
	}, nil
}

func (ur UserRepository) FindByID(ctx context.Context, id user.UserID) (*user.UserAggregate, error) {
	t, found := ur.users[id]
	if !found {
		return nil, ErrNotFound
	}

	return t, nil
}

func (tr UserRepository) Save(ctx context.Context, u *user.UserAggregate) error {
	tr.users[u.ID()] = u
	return nil
}

func (ur UserRepository) FindAll(ctx context.Context) ([]*user.UserAggregate, error) {
	var us []*user.UserAggregate
	for _, u := range ur.users {
		us = append(us, u)
	}
	return us, nil
}
