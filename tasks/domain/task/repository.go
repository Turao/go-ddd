package task

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, taskID TaskID) (*TaskAggregate, error)
	FindByProjectID(ctx context.Context, projectID ProjectID) ([]*TaskAggregate, error)
	FindByAssignedUserID(ctx context.Context, assignedUserID UserID) ([]*TaskAggregate, error)
	Save(ctx context.Context, agg *TaskAggregate) error
}
