package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/domain/task"
)

type AssignToUserCommandHandler struct {
	repository                 task.Repository
	taskAssignedEventPublisher api.TaskAssignedEventPublisher
}

func NewAssignToUserCommandHandler(repository task.Repository, taep api.TaskAssignedEventPublisher) *AssignToUserCommandHandler {
	return &AssignToUserCommandHandler{
		repository:                 repository,
		taskAssignedEventPublisher: taep,
	}
}

func (h AssignToUserCommandHandler) Handle(ctx context.Context, req application.AssignToUserCommand) error {
	agg, err := h.repository.FindByID(ctx, req.TaskID)
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, task.AssignToUserCommand{
		UserID: req.UserID,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	ie, err := api.NewTaskAssignedEvent(uuid.NewString(), req.TaskID, req.UserID)
	if err != nil {
		return err
	}

	err = h.taskAssignedEventPublisher.Publish(ctx, *ie)
	if err != nil {
		return err
	}

	return nil
}
