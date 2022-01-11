package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/turao/go-ddd/events"
	repository "github.com/turao/go-ddd/projects/adapters/sql"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/application/command"
	"github.com/turao/go-ddd/projects/application/query"
)

func main() {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			"localhost",
			5432,
			"postgres",
			"postgres",
			"postgres",
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	migrator, err := repository.NewMigrator(db, "file://projects/migrations")
	if err != nil {
		log.Fatal(err)
	}

	if err := migrator.Up(); err != nil {
		log.Fatal(err)
	}

	repo, err := repository.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	eventStore, err := events.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	app := application.App{
		Commands: application.Commands{
			CreateProject: command.NewCreateProjectCommandHandler(repo, eventStore),
			UpdateProject: command.NewUpdateProjectCommandHandler(repo, eventStore),
			DeleteProject: command.NewDeleteProjectCommandHandler(repo, eventStore),
		},
		Queries: application.Queries{
			FindProject: query.NewFindProjectQueryHandler(repo, eventStore),
		},
	}

	err = app.Commands.CreateProject.Handle(
		context.Background(),
		command.CreateProjectCommand{
			Name: "my-project",
		})

	if err != nil {
		log.Fatal(err)
	}

	err = app.Commands.UpdateProject.Handle(
		context.Background(),
		command.UpdateProjectCommand{
			ID:   "00000000-0000-0000-0000-000000000000",
			Name: "my-project-updated",
		})

	if err != nil {
		log.Fatal(err)
	}

	err = app.Commands.DeleteProject.Handle(
		context.Background(),
		command.DeleteProjectCommand{
			ID: "00000000-0000-0000-0000-000000000000",
		})

	if err != nil {
		log.Fatal(err)
	}

	evts := eventStore.Events()
	for _, evt := range evts {
		d, err := json.MarshalIndent(evt, "", " ")
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
	}

	res, err := app.Queries.FindProject.Handle(
		context.Background(),
		query.FindProjectQuery{
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
