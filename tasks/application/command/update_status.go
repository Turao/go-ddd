package command

import (
	"context"
	"log"

	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type UpdateStatusCommandHandler struct {
	repository task.Repository
	eventStore events.EventStore
}

func NewUpdateStatusCommandHandler(repository task.Repository, es events.EventStore) *UpdateStatusCommandHandler {
	return &UpdateStatusCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h UpdateStatusCommandHandler) Handle(ctx context.Context, req application.UpdateStatusCommand) error {
	t, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	ta, err := task.NewTaskAggregate(t, h.eventStore)
	if err != nil {
		return err
	}

	err = ta.UpdateStatus(req.Status)
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, *ta.Task)
	if err != nil {
		return err
	}

	// todo: should we put our integration event (task.completed) here?
	ie, err := api.NewTaskStatusUpdatedEvent(t.ID, t.Status)
	if err != nil {
		return err
	}
	log.Println("TODO: publish this task.status.updated integration event: ", ie)

	return nil
}
