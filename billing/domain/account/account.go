package account

import "errors"

type AccountID = string

var (
	ErrInvalidAccountID = errors.New("invalid account id")
)

type Account struct {
	ID      AccountID `json:"accountId"`
	User    *User     `json:"user"`
	Invoice *Invoice  `json:"invoice"`
}

func NewAccount(accountID AccountID, userID UserID, invoiceID InvoiceID) (*Account, error) {
	if accountID == "" {
		return nil, ErrInvalidAccountID
	}

	u, err := NewUser(userID)
	if err != nil {
		return nil, err
	}

	i, err := NewInvoice(invoiceID)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:      accountID,
		User:    u,
		Invoice: i,
	}, nil
}
