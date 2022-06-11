package project

import "context"

type Repository interface {
	Save(ctx context.Context, agg *ProjectAggregate) error
	FindProjectByID(ctx context.Context, id ProjectID) (*ProjectAggregate, error)
	FindAll(ctx context.Context) ([]*ProjectAggregate, error)
}
