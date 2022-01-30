package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/billing/domain/invoice"
)

type InvoiceRepository struct {
	// use the user id as primary key (i.e. only one invoice per user now)
	invoices map[invoice.UserID]*invoice.Invoice
}

var _ invoice.Repository = (*InvoiceRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewInvoiceRepository() (*InvoiceRepository, error) {
	return &InvoiceRepository{
		invoices: make(map[string]*invoice.Invoice),
	}, nil
}

func (ir InvoiceRepository) FindByUserID(ctx context.Context, userID invoice.UserID) (*invoice.Invoice, error) {
	t, found := ir.invoices[userID]
	if !found {
		return nil, ErrNotFound
	}

	return t, nil
}

func (tr InvoiceRepository) Save(ctx context.Context, p invoice.Invoice) error {
	tr.invoices[p.ID] = &p
	return nil
}

func (ir InvoiceRepository) FindAll(ctx context.Context) ([]*invoice.Invoice, error) {
	var us []*invoice.Invoice
	for _, p := range ir.invoices {
		us = append(us, p)
	}
	return us, nil
}
