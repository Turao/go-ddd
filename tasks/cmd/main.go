package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/application/command"
	"github.com/turao/go-ddd/tasks/application/query"
	"github.com/turao/go-ddd/tasks/infrastructure"
)

func main() {

	eventStore, err := events.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	tr, err := infrastructure.NewTaskRepository()
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.App{
		Commands: application.Commands{
			CreateTaskCommand:        command.NewCreateTaskCommandHandler(tr, eventStore),
			AssignToUserCommand:      command.NewAssignToUserCommandHandler(tr, eventStore),
			UnassignUserCommand:      command.NewUnassignUserCommandHandler(tr, eventStore),
			UpdateTitleCommand:       command.NewUpdateTitleCommandHandler(tr, eventStore),
			UpdateDescriptionCommand: command.NewUpdateDescriptionCommandHandler(tr, eventStore),
		},
		Queries: application.Queries{
			TasksByProjectQuery:      query.NewTaskByProjectQueryHandler(tr),
			TasksByAssignedUserQuery: query.NewTasksByAssignedUserQueryHandler(tr),
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

	// err = app.Commands.AssignToUserCommand.Handle(
	// 	context.Background(),
	// 	application.AssignToUserCommand{
	// 		TaskID: "00000000-0000-0000-0000-000000000000",
	// 		UserID: "user-id",
	// 	},
	// )
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	res, err := app.Queries.TasksByProjectQuery.Handle(
		context.Background(),
		application.TasksByProjectQuery{
			ProjectID: "projectId",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	for _, t := range res.Tasks {
		err := app.Commands.AssignToUserCommand.Handle(
			context.Background(),
			application.AssignToUserCommand{
				TaskID: t.TaskID,
				UserID: "mock-user-id",
			},
		)
		if err != nil {
			log.Fatalln(err)
		}

		// err = app.Commands.UnassignUserCommand.Handle(
		// 	context.Background(),
		// 	application.UnassignUserCommand{
		// 		TaskID: t.TaskID,
		// 	},
		// )
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		err = app.Commands.UpdateTitleCommand.Handle(
			context.Background(),
			application.UpdateTitleCommand{
				TaskID: t.TaskID,
				Title:  "this is the new title",
			},
		)
		if err != nil {
			log.Fatalln(err)
		}

		err = app.Commands.UpdateDescriptionCommand.Handle(
			context.Background(),
			application.UpdateDescriptionCommand{
				TaskID:      t.TaskID,
				Description: "this is the new description",
			},
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	res, err = app.Queries.TasksByProjectQuery.Handle(
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
	log.Println("tasks by project")
	log.Println(string(d))

	res2, err := app.Queries.TasksByAssignedUserQuery.Handle(
		context.Background(),
		application.TasksByAssignedUserQuery{
			UserID: "",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	d, err = json.MarshalIndent(res2, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("tasks by assigned user")
	log.Println(string(d))

	// dump event store
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
}
