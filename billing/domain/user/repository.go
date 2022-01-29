package user

import (
	"context"

	"github.com/turao/go-ddd/users/domain/user"
)

type Repository interface {
	FindByID(ctx context.Context, id user.UserID) (*User, error)
	Save(ctx context.Context, user User) error
	FindAll(ctx context.Context) ([]*User, error)
}
