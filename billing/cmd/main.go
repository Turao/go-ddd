package main

import (
	"context"
	"log"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/application/command"
	"github.com/turao/go-ddd/billing/infrastructure"
	"github.com/turao/go-ddd/billing/infrastructure/messaging"
	"github.com/turao/go-ddd/events/in_memory"
)

func main() {
	ur, err := infrastructure.NewAccountRepository()
	if err != nil {
		log.Fatalln(err)
	}

	es, err := in_memory.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	app := application.Application{
		Commands: application.Commands{
			CreateAccountCommand:      command.NewCreateAccountCommandHandler(ur, es),
			AddTaskToUserCommand:      command.NewAddTaskToUserCommandHandler(ur, es),
			RemoveTaskFromUserCommand: command.NewRemoveTaskFromUserCommandHandler(ur, es),
		},
		Queries: application.Queries{},
	}

	router := messaging.Router{
		CreateAccountCommandHandler:      app.Commands.CreateAccountCommand,
		AddTaskToUserCommandHandler:      app.Commands.AddTaskToUserCommand,
		RemoveTaskFromUserCommandHandler: app.Commands.RemoveTaskFromUserCommand,
	}

	err = router.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer router.Close()

	if err := router.Run(context.Background()); err != nil {
		log.Fatalln(err)
	}

}
