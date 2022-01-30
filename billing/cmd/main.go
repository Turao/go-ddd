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
	ur, err := infrastructure.NewInvoiceRepository()
	if err != nil {
		log.Fatalln(err)
	}

	es, err := in_memory.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	app := application.Application{
		Commands: application.Commands{
			CreateInvoiceCommand:         command.NewCreateInvoiceCommandHandler(ur, es),
			AddTaskToInvoiceCommand:      command.NewAddTaskToInvoiceCommandHandler(ur, es),
			RemoveTaskFromInvoiceCommand: command.NewRemoveTaskFromInvoiceCommandHandler(ur, es),
		},
		Queries: application.Queries{},
	}

	router := messaging.Router{
		CreateInvoiceCommandHandler:         app.Commands.CreateInvoiceCommand,
		AddTaskToInvoiceCommandHandler:      app.Commands.AddTaskToInvoiceCommand,
		RemoveTaskFromInvoiceCommandHandler: app.Commands.RemoveTaskFromInvoiceCommand,
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
