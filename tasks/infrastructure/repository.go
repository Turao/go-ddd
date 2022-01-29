package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/projects/domain/project"
	task "github.com/turao/go-ddd/tasks/domain"
	"github.com/turao/go-ddd/users/domain/user"
)

type TaskRepository struct {
	tasks map[task.TaskID]*task.Task
}

var _ task.Repository = (*TaskRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewTaskRepository() (*TaskRepository, error) {
	return &TaskRepository{
		tasks: make(map[string]*task.Task),
	}, nil
}

func (tr TaskRepository) FindByID(ctx context.Context, taskID task.TaskID) (*task.Task, error) {
	t, found := tr.tasks[taskID]
	if !found {
		return nil, ErrNotFound
	}
	return t, nil
}

func (tr TaskRepository) FindByProjectID(ctx context.Context, projectID project.ProjectID) ([]*task.Task, error) {
	ts := make([]*task.Task, 0)

	for _, t := range tr.tasks {
		if t.ProjectID == projectID {
			ts = append(ts, t)
		}
	}

	return ts, nil
}

func (tr TaskRepository) FindByAssignedUserID(ctx context.Context, assignedUserID user.UserID) ([]*task.Task, error) {
	ts := make([]*task.Task, 0)

	for _, t := range tr.tasks {
		if t.AssignedUser != nil {
			if *t.AssignedUser == assignedUserID {
				ts = append(ts, t)
			}
		}
	}

	return ts, nil
}

func (tr TaskRepository) Save(ctx context.Context, p task.Task) error {
	tr.tasks[p.ID] = &p
	return nil
}

// func (tr TaskRepository) FindAll(ctx context.Context) ([]*task.Task, error) {
// 	var ps []*task.Task
// 	for _, p := range pr.tasks {
// 		ps = append(ps, p)
// 	}
// 	return ps, nil
// }
