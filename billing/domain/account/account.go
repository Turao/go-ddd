package account

import (
	"errors"

	"github.com/google/uuid"
)

type AccountID = string

var (
	ErrInvalidAccountID = errors.New("invalid account id")
	ErrUserIsNil        = errors.New("user is nil")
	ErrInvoiceIsNil     = errors.New("invoice is nil")
)

type Account struct {
	ID      AccountID `json:"accountId"`
	User    *User     `json:"user"`
	Invoice *Invoice  `json:"invoice"`
}

type AccountOption = func(acc *Account) error

func WithAccountID(id string) AccountOption {
	return func(acc *Account) error {
		if id == "" {
			return ErrInvalidAccountID
		}
		acc.ID = id
		return nil
	}
}

func WithUser(user *User) AccountOption {
	return func(acc *Account) error {
		if user == nil {
			return ErrUserIsNil
		}
		acc.User = user
		return nil
	}
}

func WithInvoice(invoice *Invoice) AccountOption {
	return func(acc *Account) error {
		if invoice == nil {
			return ErrInvoiceIsNil
		}
		acc.Invoice = invoice
		return nil
	}
}

func NewAccount(opts ...AccountOption) (*Account, error) {

	acc := &Account{
		ID:      uuid.NewString(),
		User:    nil,
		Invoice: nil,
	}

	for _, opt := range opts {
		if err := opt(acc); err != nil {
			return nil, err
		}
	}

	return acc, nil
}
