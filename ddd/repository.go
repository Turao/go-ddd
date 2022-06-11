package ddd

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (Aggregate, error)
	Save(ctx context.Context, agg Aggregate) error
}
