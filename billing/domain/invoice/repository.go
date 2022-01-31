package invoice

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, invoiceID InvoiceID) (*Invoice, error)
	Save(ctx context.Context, invoice Invoice) error
	FindAll(ctx context.Context) ([]*Invoice, error)
}
