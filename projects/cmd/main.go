package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/application/command"
	"github.com/turao/go-ddd/projects/application/query"
)

func main() {

	eventStore, err := events.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	app := application.App{
		Commands: application.Commands{
			CreateProject: command.NewCreateProjectCommandHandler(eventStore),
			UpdateProject: command.NewUpdateProjectCommandHandler(eventStore),
			DeleteProject: command.NewDeleteProjectCommandHandler(eventStore),
		},
		Queries: application.Queries{
			FindProject: query.NewFindProjectQueryHandler(eventStore),
		},
	}

	err = app.Commands.CreateProject.Handle(
		context.Background(),
		application.CreateProjectCommand{
			Name: "my-project",
		})

	if err != nil {
		log.Fatal(err)
	}

	err = app.Commands.UpdateProject.Handle(
		context.Background(),
		application.UpdateProjectCommand{
			ID:   "00000000-0000-0000-0000-000000000000",
			Name: "my-project-updated",
		})

	if err != nil {
		log.Fatal(err)
	}

	err = app.Commands.DeleteProject.Handle(
		context.Background(),
		application.DeleteProjectCommand{
			ID: "00000000-0000-0000-0000-000000000000",
		})

	if err != nil {
		log.Fatal(err)
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

	res, err := app.Queries.FindProject.Handle(
		context.Background(),
		application.FindProjectQuery{
			ID: "00000000-0000-0000-0000-000000000000",
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
