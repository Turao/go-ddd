package inmemory

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/ddd"
)

type repository struct {
	aggregates map[string]ddd.Aggregate
}

var _ ddd.Repository = (*repository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewRepository() (*repository, error) {
	return &repository{
		aggregates: make(map[string]ddd.Aggregate, 0),
	}, nil
}

func (repo repository) FindByID(ctx context.Context, id string) (ddd.Aggregate, error) {
	agg, found := repo.aggregates[id]
	if !found {
		return nil, ErrNotFound
	}

	return agg, nil
}

func (repo *repository) Save(ctx context.Context, agg ddd.Aggregate) error {
	repo.aggregates[agg.ID()] = agg
	return nil
}
