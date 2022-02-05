package main

import (
	"log"

	watermillAMQP "github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/turao/go-ddd/api/amqp"
	"github.com/turao/go-ddd/events/in_memory"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/application/command"
	"github.com/turao/go-ddd/tasks/application/query"
	"github.com/turao/go-ddd/tasks/infrastructure"
)

func main() {

	eventStore, err := in_memory.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	tr, err := infrastructure.NewTaskRepository()
	if err != nil {
		log.Fatalln(err)
	}

	queue := watermillAMQP.NewDurableQueueConfig("amqp://localhost:5672")
	logger := watermill.NewStdLogger(false, false)
	publisher, err := watermillAMQP.NewPublisher(queue, logger)
	if err != nil {
		log.Fatalln(err)
	}
	defer publisher.Close()

	urep, err := amqp.NewAMQPTaskStatusUpdatedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	taep, err := amqp.NewAMQPTaskAssignedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	tuep, err := amqp.NewAMQPTaskUnassignedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.Application{
		Commands: application.Commands{
			CreateTaskCommand:        command.NewCreateTaskCommandHandler(tr, eventStore),
			AssignToUserCommand:      command.NewAssignToUserCommandHandler(tr, eventStore, taep),
			UnassignUserCommand:      command.NewUnassignUserCommandHandler(tr, eventStore, tuep),
			UpdateTitleCommand:       command.NewUpdateTitleCommandHandler(tr, eventStore),
			UpdateDescriptionCommand: command.NewUpdateDescriptionCommandHandler(tr, eventStore),
			UpdateStatusCommand:      command.NewUpdateStatusCommandHandler(tr, eventStore, urep),
		},
		Queries: application.Queries{
			TasksByProjectQuery:      query.NewTaskByProjectQueryHandler(tr),
			TasksByAssignedUserQuery: query.NewTasksByAssignedUserQueryHandler(tr),
		},
	}

	server, err := infrastructure.NewServer(&infrastructure.Application{
		Delegate: app,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
