package main

import (
	"context"
	"encoding/json"
	"log"

	watermillAMQP "github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/turao/go-ddd/api/amqp"
	"github.com/turao/go-ddd/events/in_memory"
	"github.com/turao/go-ddd/tasks/application"
	"github.com/turao/go-ddd/tasks/application/command"
	"github.com/turao/go-ddd/tasks/application/query"
	"github.com/turao/go-ddd/tasks/infrastructure"
)

func main() {

	eventStore, err := in_memory.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	tr, err := infrastructure.NewTaskRepository()
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

	urep, err := amqp.NewAMQPTaskStatusUpdatedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	taep, err := amqp.NewAMQPTaskAssignedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	tuep, err := amqp.NewAMQPTaskUnassignedEventPublisher(publisher)
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.App{
		Commands: application.Commands{
			CreateTaskCommand:        command.NewCreateTaskCommandHandler(tr, eventStore),
			AssignToUserCommand:      command.NewAssignToUserCommandHandler(tr, eventStore, taep),
			UnassignUserCommand:      command.NewUnassignUserCommandHandler(tr, eventStore, tuep),
			UpdateTitleCommand:       command.NewUpdateTitleCommandHandler(tr, eventStore),
			UpdateDescriptionCommand: command.NewUpdateDescriptionCommandHandler(tr, eventStore),
			UpdateStatusCommand:      command.NewUpdateStatusCommandHandler(tr, eventStore, urep),
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

	// err = app.Commands.CreateTaskCommand.Handle(
	// 	context.Background(),
	// 	application.CreateTaskCommand{
	// 		ProjectID:   "projectId",
	// 		Title:       "task-title",
	// 		Description: "task-description",
	// 	},
	// )
	// if err != nil {
	// 	log.Fatalln(err)
	// }

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
				UserID: "00000000-0000-0000-0000-000000000000",
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

		err = app.Commands.UpdateStatusCommand.Handle(
			context.Background(),
			application.UpdateStatusCommand{
				TaskID: t.TaskID,
				Status: "completed",
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
