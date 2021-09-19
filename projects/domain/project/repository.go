package project

import "context"



type WriteRepository interface {
	Create(ctx context.Context, p Project) error
	Update(ctx context.Context, p Project) error
	Delete(ctx context.Context, id ProjectID) error
}

type ReadRepository interface {
	FindProjectByID(ctx context.Context, id ProjectID) (*Project, error)
}
