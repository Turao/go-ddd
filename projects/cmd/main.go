package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events/in_memory"
	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/application/command"
	"github.com/turao/go-ddd/projects/application/query"
	"github.com/turao/go-ddd/projects/infrastructure"
)

func main() {

	eventStore, err := in_memory.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	pr, err := infrastructure.NewProjectRepository()
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.Application{
		Commands: application.Commands{
			CreateProject: command.NewCreateProjectCommandHandler(pr, eventStore),
			UpdateProject: command.NewUpdateProjectCommandHandler(pr, eventStore),
			DeleteProject: command.NewDeleteProjectCommandHandler(pr, eventStore),
		},
		Queries: application.Queries{
			FindProject:  query.NewFindProjectQueryHandler(pr),
			ListProjects: query.NewListProjectsQueryHandler(pr),
		},
	}

	err = app.Commands.CreateProject.Handle(
		context.Background(),
		application.CreateProjectCommand{
			Name:      "my-project",
			CreatedBy: uuid.NewString(),
		})

	if err != nil {
		log.Fatal(err)
	}

	// list all
	res, err := app.Queries.ListProjects.Handle(
		context.Background(),
		application.ListProjectsQuery{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	d, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("listing all projects")
	log.Println(string(d))

	// iterate and find by id
	for _, p := range res.Projects {
		log.Println("searching for project: ", p.ID)
		res2, err := app.Queries.FindProject.Handle(
			context.Background(),
			application.FindProjectQuery{
				ID: p.ID,
			},
		)
		if err != nil {
			log.Fatalln(err)
		}

		d, err := json.MarshalIndent(res2, "", " ")
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
	}

	// err = app.Commands.UpdateProject.Handle(
	// 	context.Background(),
	// 	application.UpdateProjectCommand{
	// 		ID:   "00000000-0000-0000-0000-000000000000",
	// 		Name: "my-project-updated",
	// 	})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = app.Commands.DeleteProject.Handle(
	// 	context.Background(),
	// 	application.DeleteProjectCommand{
	// 		ID: "00000000-0000-0000-0000-000000000000",
	// 	})

	// if err != nil {
	// 	log.Fatal(err)
	// }

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

	// res, err := app.Queries.FindProject.Handle(
	// 	context.Background(),
	// 	application.FindProjectQuery{
	// 		ID: "00000000-0000-0000-0000-000000000000",
	// 	},
	// )
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// d, err := json.MarshalIndent(res, "", " ")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(string(d))

	server, err := infrastructure.NewServer(&infrastructure.Application{
		Delegate: app,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}
