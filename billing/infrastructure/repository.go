package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/billing/domain/account"
)

type AccountRepository struct {
	accounts map[account.AccountID]*account.Account
}

var _ account.Repository = (*AccountRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewAccountRepository() (*AccountRepository, error) {
	return &AccountRepository{
		accounts: make(map[string]*account.Account),
	}, nil
}

func (ir AccountRepository) FindByID(ctx context.Context, accountID account.AccountID) (*account.Account, error) {
	t, found := ir.accounts[accountID]
	if !found {
		return nil, ErrNotFound
	}

	return t, nil
}

func (tr AccountRepository) Save(ctx context.Context, p account.Account) error {
	tr.accounts[p.ID] = &p
	return nil
}

func (ir AccountRepository) FindAll(ctx context.Context) ([]*account.Account, error) {
	var us []*account.Account
	for _, p := range ir.accounts {
		us = append(us, p)
	}
	return us, nil
}
