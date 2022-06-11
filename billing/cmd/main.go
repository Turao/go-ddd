package main

import (
	"context"
	"log"
	"sync"

	"github.com/turao/go-ddd/billing/application"
	"github.com/turao/go-ddd/billing/application/command"
	"github.com/turao/go-ddd/billing/application/query"
	"github.com/turao/go-ddd/billing/infrastructure/messaging"
	"github.com/turao/go-ddd/billing/infrastructure/rest"
	"github.com/turao/go-ddd/ddd/inmemory"
	"github.com/turao/go-ddd/events/in_memory"
)

func main() {
	repo, err := inmemory.NewRepository()
	if err != nil {
		log.Fatalln(err)
	}

	es, err := in_memory.NewInMemoryStore()
	if err != nil {
		log.Fatalln(err)
	}

	app := &application.Application{
		Commands: application.Commands{
			CreateAccountCommand:      command.NewCreateAccountCommandHandler(repo, es),
			AddTaskToUserCommand:      command.NewAddTaskToUserCommandHandler(repo, es),
			RemoveTaskFromUserCommand: command.NewRemoveTaskFromUserCommandHandler(repo, es),
		},
		Queries: application.Queries{
			GetAccountDetails: query.NewGetAccountDetailsQueryHandler(repo, es),
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	// setup messaging integration
	go func() {
		defer wg.Done()

		router := messaging.Router{
			CreateAccountCommandHandler:      app.Commands.CreateAccountCommand,
			AddTaskToUserCommandHandler:      app.Commands.AddTaskToUserCommand,
			RemoveTaskFromUserCommandHandler: app.Commands.RemoveTaskFromUserCommand,
		}

		err = router.Init()
		if err != nil {
			log.Fatalln(err)
		}
		defer router.Close()

		if err := router.Run(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}()

	// setup HTTP server
	go func() {
		wg.Done()

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
	}()

	// wait for all adapters to close
	wg.Wait()

}
