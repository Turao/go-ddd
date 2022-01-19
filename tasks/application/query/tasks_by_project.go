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
	ts, err := h.repository.FindByProjectID(ctx, req.ProjectID)
	if err != nil {
		return nil, err
	}

	tsDTO := make([]application.Task, 0)
	for _, t := range ts {
		assignedTo := ""
		if t.AssignedUser != nil {
			assignedTo = *t.AssignedUser
		}

		tsDTO = append(tsDTO, application.Task{
			TaskID:     t.ID,
			AssignedTo: assignedTo,
		})
	}

	return &application.TasksByProjectResponse{
		ProjectID: req.ProjectID,
		Tasks:     tsDTO,
	}, nil
}
