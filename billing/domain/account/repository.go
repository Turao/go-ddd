package account

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, accountID string) (*AccountAggregate, error)
	Save(ctx context.Context, account *AccountAggregate) error
	FindAll(ctx context.Context) ([]*AccountAggregate, error)
}
