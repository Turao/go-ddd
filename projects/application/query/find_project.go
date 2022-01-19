package query

import (
	"context"

	"github.com/turao/go-ddd/projects/application"
	"github.com/turao/go-ddd/projects/domain/project"
)

type FindProjectHandler struct {
	repository project.Repository
}

func NewFindProjectQueryHandler(repository project.Repository) *FindProjectHandler {
	return &FindProjectHandler{
		repository: repository,
	}
}

func (h *FindProjectHandler) Handle(ctx context.Context, req application.FindProjectQuery) (*application.FindProjectResponse, error) {
	p, err := h.repository.FindProjectByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &application.FindProjectResponse{
		ID:        p.ID,
		Name:      p.Name,
		CreatedBy: p.CreatedBy,
		CreatedAt: p.CreatedAt,
		Active:    p.Active,
	}, nil
}
