package main

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	watermillKafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/turao/go-ddd/api/kafka"
	"github.com/turao/go-ddd/users/application"
	"github.com/turao/go-ddd/users/application/command"
	"github.com/turao/go-ddd/users/application/query"
	"github.com/turao/go-ddd/users/infrastructure"
	"github.com/turao/go-ddd/users/infrastructure/rest"
)

func main() {
	ur, err := infrastructure.NewUserRepository()
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

	urep, err := kafka.NewUserRegisteredEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.Application{
		Commands: application.Commands{
			RegisterUserCommand: command.NewRegisterUserHandler(ur, urep),
		},
		Queries: application.Queries{
			ListUsersQuery: query.NewListUsersQueryHandler(ur),
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
