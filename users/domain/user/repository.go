package user

import "context"

type Repository interface {
	FindByID(ctx context.Context, id UserID) (*UserAggregate, error)
	Save(ctx context.Context, user *UserAggregate) error
	FindAll(ctx context.Context) ([]*UserAggregate, error)
}
