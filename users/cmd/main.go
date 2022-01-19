package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/turao/go-ddd/events"
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

	es, err := events.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	app := application.Application{
		Commands: application.Commands{
			RegisterUserCommand: command.NewRegisterUserHandler(ur, es),
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

}
