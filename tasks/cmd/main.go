package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/application/command"
	"github.com/turao/go-ddd/tasks/application/query"
)

func main() {

	eventStore, err := events.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.App{
		Commands: application.Commands{
			CreateTaskCommand: command.NewCreateTaskCommandHandler(eventStore),
		},
		Queries: application.Queries{
			TasksByProjectQuery: query.NewTaskByProjectQueryHandler(eventStore),
		},
	}

	err = app.Commands.CreateTaskCommand.Handle(
		context.Background(),
		application.CreateTaskCommand{
			ProjectID:   "projectId",
			Title:       "task-title",
			Description: "task-description",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Commands.CreateTaskCommand.Handle(
		context.Background(),
		application.CreateTaskCommand{
			ProjectID:   "projectId",
			Title:       "task-title",
			Description: "task-description",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	evts, err := eventStore.Events(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for _, evt := range evts {
		d, err := json.MarshalIndent(evt, "", " ")
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
	}

	res, err := app.Queries.TasksByProjectQuery.Handle(
		context.Background(),
		application.TasksByProjectQuery{
			ProjectID: "projectId",
		},
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
