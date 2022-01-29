package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/billing/domain/user"
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

func (ur UserRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	var us []*user.User
	for _, p := range ur.users {
		us = append(us, p)
	}
	return us, nil
}
