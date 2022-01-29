package task

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, taskID TaskID) (*Task, error)
	FindByProjectID(ctx context.Context, projectID ProjectID) ([]*Task, error)
	FindByAssignedUserID(ctx context.Context, assignedUserID UserID) ([]*Task, error)
	Save(ctx context.Context, task Task) error
}
