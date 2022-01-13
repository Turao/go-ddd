package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type AddTaskCommand struct {
	eventStore events.EventStore
}

func NewAddTaskCommandHandler(es events.EventStore) *AddTaskCommand {
	return &AddTaskCommand{
		eventStore: es,
	}
}

func (atc *AddTaskCommand) Handle(ctx context.Context, req application.AddTaskCommand) error {
	evt, err := project.NewTaskAddedEvent(req.ID, req.TaskID)
	if err != nil {
		return nil
	}

	return atc.eventStore.Push(ctx, *evt)
}
