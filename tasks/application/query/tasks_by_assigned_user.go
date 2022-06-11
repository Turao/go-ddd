package query

import (
	"context"

	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type TasksByAssignedUserQueryHandler struct {
	repository task.Repository
}

func NewTasksByAssignedUserQueryHandler(repository task.Repository) *TasksByAssignedUserQueryHandler {
	return &TasksByAssignedUserQueryHandler{
		repository: repository,
	}
}

func (h TasksByAssignedUserQueryHandler) Handle(ctx context.Context, req application.TasksByAssignedUserQuery) (*application.TasksByAssignedUserResponse, error) {
	tasks, err := h.repository.FindByAssignedUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	tasksDTO := make([]application.Task, 0)
	for _, task := range tasks {
		assignedUser := ""
		if task.Task.AssignedUser != nil {
			assignedUser = *task.Task.AssignedUser
		}

		tasksDTO = append(tasksDTO, application.Task{
			TaskID:     task.Task.ID,
			AssignedTo: assignedUser,
			Status:     task.Task.Status,
		})
	}

	return &application.TasksByAssignedUserResponse{
		UserID: req.UserID,
		Tasks:  tasksDTO,
	}, nil
}
