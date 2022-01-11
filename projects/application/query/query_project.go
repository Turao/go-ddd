package query

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type FindProjectQuery struct {
	ID string `json:"id"`
}

type FindProjectResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FindProjectHandler struct {
	repo project.Repository
}

func NewFindProjectQueryHandler(repo project.Repository) *FindProjectHandler {
	return &FindProjectHandler{
		repo: repo,
	}
}

func (h *FindProjectHandler) Handle(ctx context.Context, req FindProjectQuery) (*FindProjectResponse, error) {
	p, err := h.repo.FindProjectByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &FindProjectResponse{
		ID:   p.ID,
		Name: p.Name,
	}, nil
}
