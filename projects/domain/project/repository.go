package project

import "context"

type Repository interface {
	Save(ctx context.Context, p Project) error
	FindProjectByID(ctx context.Context, id ProjectID) (*Project, error)
}
