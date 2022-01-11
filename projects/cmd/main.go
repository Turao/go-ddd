package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	repository "github.com/turao/go-ddd/projects/adapters/sql"
	"github.com/turao/go-ddd/projects/domain/project"
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

	// repo, err := repository.NewRepository(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// app := application.App{
	// 	Commands: application.Commands{
	// 		CreateProject: command.NewCreateProjectCommandHandler(repo),
	// 		UpdateProject: command.NewUpdateProjectCommandHandler(repo),
	// 		DeleteProject: command.NewDeleteProjectCommandHandler(repo),
	// 	},
	// 	Queries: application.Queries{
	// 		FindProject: query.NewFindProjectQueryHandler(repo),
	// 	},
	// }

	// err = app.Commands.CreateProject.Handle(
	// 	context.Background(),
	// 	command.CreateProjectCommand{
	// 		Title: "my-title",
	// 	})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = app.Commands.UpdateProject.Handle(
	// 	context.Background(),
	// 	command.UpdateProjectCommand{
	// 		ID:    "00000000-0000-0000-0000-000000000000",
	// 		Title: "my-title",
	// 	})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	agg := &project.ProjectAggregate{
		Project: nil,
	}

	id := uuid.NewString()
	createEvent, err := project.NewProjectCreatedEvent(id, "my-project")
	if err != nil {
		log.Fatal(err)
	}

	err = agg.Handle(*createEvent)
	if err != nil {
		log.Fatal(err)
	}

	deleteEvent, err := project.NewProjectDeletedEvent(id)
	if err != nil {
		log.Fatal(err)
	}

	err = agg.Handle(*deleteEvent)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(agg)
}
