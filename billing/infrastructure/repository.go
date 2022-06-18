package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/billing/domain/account"
)

type AccountRepository struct {
	accounts map[account.AccountID]*account.AccountAggregate
}

var _ account.Repository = (*AccountRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewAccountRepository() (*AccountRepository, error) {
	return &AccountRepository{
		accounts: make(map[string]*account.AccountAggregate),
	}, nil
}

func (ir AccountRepository) FindByID(ctx context.Context, accountID account.AccountID) (*account.AccountAggregate, error) {
	acc, found := ir.accounts[accountID]
	if !found {
		return nil, ErrNotFound
	}

	return acc, nil
}

func (tr *AccountRepository) Save(ctx context.Context, acc *account.AccountAggregate) error {
	tr.accounts[acc.ID()] = acc
	return nil
}

func (ir AccountRepository) FindAll(ctx context.Context) ([]*account.AccountAggregate, error) {
	var accs []*account.AccountAggregate
	for _, acc := range ir.accounts {
		accs = append(accs, acc)
	}
	return accs, nil
}
