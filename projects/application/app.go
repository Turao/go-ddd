package application

import (
	"github.com/turao/go-ddd/projects/application/command"
	"github.com/turao/go-ddd/projects/application/query"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateProject *command.CreateProjectHandler
	UpdateProject *command.UpdateProjectHandler
	DeleteProject *command.DeleteProjectHandler
}

type Queries struct {
	FindProject query.FindProjectHandler
}
