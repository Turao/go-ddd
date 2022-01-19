package query

import (
	"context"

	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type ListProjectsQueryHandler struct {
	repository project.Repository
}

func NewListProjectsQueryHandler(repository project.Repository) *ListProjectsQueryHandler {
	return &ListProjectsQueryHandler{
		repository: repository,
	}
}

func (q ListProjectsQueryHandler) Handle(ctx context.Context, req application.ListProjectsQuery) (*application.ListProjectsResponse, error) {
	ps, err := q.repository.FindAll(ctx)
	if err != nil {
		return &application.ListProjectsResponse{
			Projects: make([]application.Project, 0),
		}, err
	}

	var psDTOs []application.Project
	for _, p := range ps {
		psDTOs = append(psDTOs, application.Project{
			ID:        p.ID,
			Name:      p.Name,
			CreatedBy: p.CreatedBy,
			CreatedAt: p.CreatedAt,
			Active:    p.Active,
		})
	}
	return &application.ListProjectsResponse{
		Projects: psDTOs,
	}, nil
}
