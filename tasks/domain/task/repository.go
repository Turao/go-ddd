package task

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
	"github.com/turao/go-ddd/users/domain/user"
)

type Repository interface {
	FindByID(ctx context.Context, taskID TaskID) (*Task, error)
	FindByProjectID(ctx context.Context, projectID project.ProjectID) ([]*Task, error)
	FindByAssignedUserID(ctx context.Context, assignedUserID user.UserID) ([]*Task, error)
	Save(ctx context.Context, task Task) error
}
