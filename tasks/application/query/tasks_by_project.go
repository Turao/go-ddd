package query

import (
	"context"
	"log"

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
	log.Println("querying tasks by project id", req)

	ts, err := h.repository.FindByProjectID(ctx, req.ProjectID)
	if err != nil {
		return nil, err
	}

	tsDTO := make([]application.Task, 0)
	for _, t := range ts {
		tsDTO = append(tsDTO, application.Task{
			TaskID: t.ID,
		})
	}

	return &application.TasksByProjectResponse{
		ProjectID: req.ProjectID,
		Tasks:     tsDTO,
	}, nil
}
