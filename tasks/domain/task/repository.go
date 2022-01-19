package task

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type Repository interface {
	FindByProjectID(ctx context.Context, projectID project.ProjectID) ([]*Task, error)
	Save(ctx context.Context, task Task) error
}
