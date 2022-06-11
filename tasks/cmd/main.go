package main

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	watermillKafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/turao/go-ddd/api/kafka"
	"github.com/turao/go-ddd/events/inmemory"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/application/command"
	"github.com/turao/go-ddd/tasks/application/query"
	"github.com/turao/go-ddd/tasks/infrastructure"
)

func main() {

	eventStore, err := inmemory.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

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
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
