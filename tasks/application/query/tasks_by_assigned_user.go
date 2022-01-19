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
	ts, err := h.repository.FindByAssignedUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	tsDTO := make([]application.Task, 0)
	for _, t := range ts {
		assignedUser := ""
		if t.AssignedUser != nil {
			assignedUser = *t.AssignedUser
		}

		tsDTO = append(tsDTO, application.Task{
			TaskID:     t.ID,
			AssignedTo: assignedUser,
		})
	}

	return &application.TasksByAssignedUserResponse{
		UserID: req.UserID,
		Tasks:  tsDTO,
	}, nil
}
