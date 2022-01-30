package invoice

import (
	"context"

	"github.com/turao/go-ddd/users/domain/user"
)

type Repository interface {
	FindByUserID(ctx context.Context, userID user.UserID) (*Invoice, error)
	Save(ctx context.Context, invoice Invoice) error
	FindAll(ctx context.Context) ([]*Invoice, error)
}
