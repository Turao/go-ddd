package main

import (
	"log"

	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/application/command"
	"github.com/turao/go-ddd/projects/application/query"
	"github.com/turao/go-ddd/projects/infrastructure"
	"github.com/turao/go-ddd/projects/infrastructure/rest"
)

func main() {
	pr, err := infrastructure.NewProjectRepository()
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.Application{
		Commands: application.Commands{
			CreateProject: command.NewCreateProjectCommandHandler(pr),
			UpdateProject: command.NewUpdateProjectCommandHandler(pr),
			DeleteProject: command.NewDeleteProjectCommandHandler(pr),
		},
		Queries: application.Queries{
			FindProject:  query.NewFindProjectQueryHandler(pr),
			ListProjects: query.NewListProjectsQueryHandler(pr),
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
