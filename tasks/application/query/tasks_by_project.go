package query

import (
	"context"
	"log"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
)

type TasksByProjectQueryHandler struct {
	eventStore events.EventStore
}

func NewTaskByProjectQueryHandler(es events.EventStore) *TasksByProjectQueryHandler {
	return &TasksByProjectQueryHandler{
		eventStore: es,
	}
}

func (tbp TasksByProjectQueryHandler) Handle(
	ctx context.Context,
	req application.TasksByProjectQuery,
) (*application.TasksByProjectResponse, error) {
	log.Println("querying tasks by project id", req)
	return nil, nil
}
