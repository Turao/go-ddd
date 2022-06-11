package main

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	watermillKafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/turao/go-ddd/api/kafka"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/application/command"
	"github.com/turao/go-ddd/tasks/application/query"
	"github.com/turao/go-ddd/tasks/infrastructure"
	"github.com/turao/go-ddd/tasks/infrastructure/rest"
)

func main() {
	tr, err := infrastructure.NewTaskRepository()
	if err != nil {
		log.Fatalln(err)
	}

	queue := watermillKafka.PublisherConfig{
		Brokers:   []string{"localhost:29092"},
		Marshaler: watermillKafka.DefaultMarshaler{},
	}
	logger := watermill.NewStdLogger(false, false)
	publisher, err := watermillKafka.NewPublisher(queue, logger)
	if err != nil {
		log.Fatalln(err)
	}
	defer publisher.Close()

	urep, err := kafka.NewTaskStatusUpdatedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	taep, err := kafka.NewTaskAssignedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	tuep, err := kafka.NewTaskUnassignedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.Application{
		Commands: application.Commands{
			CreateTaskCommand:        command.NewCreateTaskCommandHandler(tr),
			AssignToUserCommand:      command.NewAssignToUserCommandHandler(tr, taep),
			UnassignUserCommand:      command.NewUnassignUserCommandHandler(tr, tuep),
			UpdateTitleCommand:       command.NewUpdateTitleCommandHandler(tr),
			UpdateDescriptionCommand: command.NewUpdateDescriptionCommandHandler(tr),
			UpdateStatusCommand:      command.NewUpdateStatusCommandHandler(tr, urep),
		},
		Queries: application.Queries{
			TasksByProjectQuery:      query.NewTaskByProjectQueryHandler(tr),
			TasksByAssignedUserQuery: query.NewTasksByAssignedUserQueryHandler(tr),
		},
	}

	server, err := rest.NewServer(&rest.Application{
		Delegate: app,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
