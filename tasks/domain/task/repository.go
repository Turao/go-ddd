package task

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type Repository interface {
	FindByID(ctx context.Context, taskID TaskID) (*Task, error)
	FindByProjectID(ctx context.Context, projectID project.ProjectID) ([]*Task, error)
	Save(ctx context.Context, task Task) error
}
