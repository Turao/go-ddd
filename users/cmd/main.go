package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	watermillAMQP "github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/turao/go-ddd/api/amqp"
	"github.com/turao/go-ddd/events/in_memory"
	"github.com/turao/go-ddd/users/application"
	"github.com/turao/go-ddd/users/application/command"
	"github.com/turao/go-ddd/users/application/query"
	"github.com/turao/go-ddd/users/infrastructure"
)

func main() {
	ur, err := infrastructure.NewUserRepository()
	if err != nil {
		log.Fatalln(err)
	}

	es, err := in_memory.NewInMemoryStore()
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

	urep, err := amqp.NewAMQPUserRegisteredEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	app := application.Application{
		Commands: application.Commands{
			RegisterUserCommand: command.NewRegisterUserHandler(ur, es, urep),
		},
		Queries: application.Queries{
			ListUsersQuery: query.NewListUsersQueryHandler(ur),
		},
	}

	err = app.Commands.RegisterUserCommand.Handle(
		context.Background(),
		application.RegisterUserCommand{
			Username: "turao",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Commands.RegisterUserCommand.Handle(
		context.Background(),
		application.RegisterUserCommand{
			Username: "lenz",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := app.Queries.ListUsersQuery.Handle(
		context.Background(),
		application.ListUsersQuery{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	d, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(d))

	subscriber, err := watermillAMQP.NewSubscriber(queue, logger)
	ures, err := amqp.NewUserRegisteredEventSubscriber(subscriber)
	if err != nil {
		log.Fatalln(err)
	}
	events, err := ures.Subscribe(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for event := range events {
			log.Println("received event:", event)
		}
	}()

	time.Sleep(10 * time.Second)

}
