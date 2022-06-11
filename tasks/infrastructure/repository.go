package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/tasks/domain/task"
)

type TaskRepository struct {
	aggregates map[task.TaskID]*task.TaskAggregate
}

var _ task.Repository = (*TaskRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewTaskRepository() (*TaskRepository, error) {
	return &TaskRepository{
		aggregates: make(map[string]*task.TaskAggregate),
	}, nil
}

func (tr TaskRepository) FindByID(ctx context.Context, taskID task.TaskID) (*task.TaskAggregate, error) {
	t, found := tr.aggregates[taskID]
	if !found {
		return nil, ErrNotFound
	}
	return t, nil
}

func (tr TaskRepository) FindByProjectID(ctx context.Context, projectID task.ProjectID) ([]*task.TaskAggregate, error) {
	ts := make([]*task.TaskAggregate, 0)

	for _, agg := range tr.aggregates {
		if agg.Task.ProjectID == projectID {
			ts = append(ts, agg)
		}
	}

	return ts, nil
}

func (tr TaskRepository) FindByAssignedUserID(ctx context.Context, assignedUserID task.UserID) ([]*task.TaskAggregate, error) {
	ts := make([]*task.TaskAggregate, 0)

	for _, agg := range tr.aggregates {
		if agg.Task.AssignedUser != nil {
			if *agg.Task.AssignedUser == assignedUserID {
				ts = append(ts, agg)
			}
		}
	}

	return ts, nil
}

func (tr TaskRepository) Save(ctx context.Context, agg *task.TaskAggregate) error {
	tr.aggregates[agg.ID()] = agg
	return nil
}
