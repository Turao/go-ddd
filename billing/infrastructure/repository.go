package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/billing/domain/invoice"
)

type InvoiceRepository struct {
	invoices map[invoice.InvoiceID]*invoice.Invoice
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

func (ir InvoiceRepository) FindByID(ctx context.Context, invoiceID invoice.InvoiceID) (*invoice.Invoice, error) {
	t, found := ir.invoices[invoiceID]
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
