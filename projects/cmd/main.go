package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

	migrator, err := repository.NewMigrator(db, "file://migrations")
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

	app := application.App{
		Commands: application.Commands{
			CreateProject: command.NewCreateProjectCommandHandler(repo),
			UpdateProject: command.NewUpdateProjectCommandHandler(repo),
			DeleteProject: command.NewDeleteProjectCommandHandler(repo),
		},
		Queries: application.Queries{
			FindProject: query.NewFindProjectQueryHandler(repo),
		},
	}

	err = app.Commands.CreateProject.Handle(
		context.Background(),
		command.CreateProjectCommand{
			Title: "my-title",
		})

	if err != nil {
		log.Fatal(err)
	}

	err = app.Commands.UpdateProject.Handle(
		context.Background(),
		command.UpdateProjectCommand{
			ID:    "00000000-0000-0000-0000-000000000000",
			Title: "my-title",
		})

	if err != nil {
		log.Fatal(err)
	}
}
