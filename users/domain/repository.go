package domain

import "context"

type Repository interface {
	FindByID(ctx context.Context, id UserID) (*User, error)
	Save(ctx context.Context, user User) error
	FindAll(ctx context.Context) ([]*User, error)
}
