package query

import (
	"context"

	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type TasksByProjectQueryHandler struct {
	repository task.Repository
}

func NewTaskByProjectQueryHandler(repository task.Repository) *TasksByProjectQueryHandler {
	return &TasksByProjectQueryHandler{
		repository: repository,
	}
}

func (h TasksByProjectQueryHandler) Handle(
	ctx context.Context,
	req application.TasksByProjectQuery,
) (*application.TasksByProjectResponse, error) {
	tasks, err := h.repository.FindByProjectID(ctx, req.ProjectID)
	if err != nil {
		return nil, err
	}

	tasksDTO := make([]application.Task, 0)
	for _, task := range tasks {
		assignedTo := ""
		if task.Task.AssignedUser != nil {
			assignedTo = *task.Task.AssignedUser
		}

		tasksDTO = append(tasksDTO, application.Task{
			TaskID:     task.Task.ID,
			AssignedTo: assignedTo,
			Status:     task.Task.Status,
		})
	}

	return &application.TasksByProjectResponse{
		ProjectID: req.ProjectID,
		Tasks:     tasksDTO,
	}, nil
}
