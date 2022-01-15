package query

import (
	"context"
	"log"
	"reflect"

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

func (h TasksByProjectQueryHandler) Handle(
	ctx context.Context,
	req application.TasksByProjectQuery,
) (*application.TasksByProjectResponse, error) {
	log.Println("querying tasks by project id", req)

	evts, err := h.eventStore.Events(context.Background())
	if err != nil {
		return nil, err
	}

	log.Println("events", evts)

	for _, evt := range evts {
		log.Println(reflect.TypeOf(evt))

		devt := evt.(events.DomainEvent)
		if devt.AggregateID() == req.ProjectID {
			log.Println(evt)
		}
	}

	return nil, nil
}
