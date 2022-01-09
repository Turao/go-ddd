package command

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type CreateProjectCommand struct {
	Title string `json:"title"`
}

type CreateProjectHandler struct {
	repo project.Repository
}

func NewCreateProjectCommandHandler(repo project.Repository) *CreateProjectHandler {
	return &CreateProjectHandler{
		repo: repo,
	}
}

func (h *CreateProjectHandler) Handle(ctx context.Context, req CreateProjectCommand) error {
	p, err := project.From(req.Title)
	if err != nil {
		return err
	}

	if err := h.repo.Save(ctx, *p); err != nil {
		return err
	}

	return nil
}
