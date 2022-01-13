package command

import (
	"context"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type RemoveTaskCommand struct {
	eventStore events.EventStore
}

func NewRemoveTaskCommandHandler(es events.EventStore) *RemoveTaskCommand {
	return &RemoveTaskCommand{
		eventStore: es,
	}
}

func (atc *RemoveTaskCommand) Handle(ctx context.Context, req application.RemoveTaskCommand) error {
	evt, err := project.NewTaskRemovedEvent(req.ID, req.TaskID)
	if err != nil {
		return nil
	}

	return atc.eventStore.Push(ctx, *evt)
}
