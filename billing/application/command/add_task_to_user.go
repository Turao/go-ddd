package command

import (
	"context"
	"log"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/domain/account"
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/ddd/eventsource"
)

type AddTaskToUserCommandHandler struct {
	repository ddd.Repository
	eventStore ddd.DomainEventStore
}

var _ application.AddTaskToUserCommandHandler = (*AddTaskToUserCommandHandler)(nil)

func NewAddTaskToUserCommandHandler(repository ddd.Repository, es ddd.DomainEventStore) *AddTaskToUserCommandHandler {
	return &AddTaskToUserCommandHandler{
		repository: repository,
		eventStore: es,
	}
}

func (h AddTaskToUserCommandHandler) Handle(ctx context.Context, req application.AddTaskToUserCommand) error {
	defer func() {
		data, _ := h.eventStore.MarshalJSON()
		log.Println("after creating")
		log.Println(string(data))
	}()

	agg, err := account.NewAggregate(account.AccountEventsFactory{})
	if err != nil {
		return err
	}

	root, err := eventsource.NewAggregate(agg, h.eventStore)
	if err != nil {
		return err
	}

	err = root.ReplayEvents()
	if err != nil {
		return err
	}

	_, err = agg.HandleCommand(ctx, account.AddTaskToUserCommand{
		TaskID: req.TaskID,
	})
	if err != nil {
		return err
	}

	err = h.repository.Save(ctx, agg)
	if err != nil {
		return err
	}

	return nil
}
