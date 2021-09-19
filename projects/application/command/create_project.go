package command

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type CreateProjectRequest struct {
	Title string `json:"title"`
}

type CreateProjectHandler struct {
	repo project.WriteRepository
}

func NewCreateProjectCommandHandler(repo project.WriteRepository) *CreateProjectHandler {
	return &CreateProjectHandler{
		repo: repo,
	}
}

func (h *CreateProjectHandler) Handle(ctx context.Context, req CreateProjectRequest) error {
	p, err := project.From(req.Title)
	if err != nil {
		return err
	}

	if err := h.repo.Create(ctx, *p); err != nil {
		return err
	}

	return nil
}
