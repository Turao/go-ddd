package application

import (
	"context"
	"time"

	"github.com/turao/go-ddd/projects/domain/project"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateProject CreateProjectHandler
	UpdateProject UpdateProjectHandler
	DeleteProject DeleteProjectHandler
}

type Queries struct {
	FindProject  FindProjectQueryHandler
	ListProjects ListProjectsQueryHandler
}

// -- Commands --
type CreateProjectCommand struct {
	Name      string `json:"name"`
	CreatedBy string `json:"createdBy"`
}

type CreateProjectHandler interface {
	Handle(ctx context.Context, req CreateProjectCommand) error
}

type DeleteProjectCommand struct {
	ID project.ProjectID `json:"id"`
}

type DeleteProjectHandler interface {
	Handle(ctx context.Context, req DeleteProjectCommand) error
}

type UpdateProjectCommand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateProjectHandler interface {
	Handle(ctx context.Context, req UpdateProjectCommand) error
}

// -- Queries --

type FindProjectQuery struct {
	ID string `json:"id"`
}

type FindProjectResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`

	Active bool `json:"active"`
}

type FindProjectQueryHandler interface {
	Handle(ctx context.Context, req FindProjectQuery) (*FindProjectResponse, error)
}

type ListProjectsQuery struct{}

type ListProjectsResponse struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`

	Active bool `json:"active"`
}

type ListProjectsQueryHandler interface {
	Handle(ctx context.Context, req ListProjectsQuery) (*ListProjectsResponse, error)
}
