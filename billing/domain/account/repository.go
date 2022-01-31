package account

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, accountID AccountID) (*Account, error)
	Save(ctx context.Context, account Account) error
	FindAll(ctx context.Context) ([]*Account, error)
}
