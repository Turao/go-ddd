package query

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type FindProjectRequest struct {
	ID string `json:"id"`
}

type FindProjectResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type FindProjectHandler struct {
	repo project.ReadRepository
}

func NewFindProjectCommandHandler(repo project.ReadRepository) *FindProjectHandler {
	return &FindProjectHandler{
		repo: repo,
	}
}

func (h *FindProjectHandler) Handle(ctx context.Context, req FindProjectRequest) (*FindProjectResponse, error) {
	p, err := h.repo.FindProjectByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &FindProjectResponse{
		ID:    p.ID,
		Title: p.Title,
	}, nil
}
