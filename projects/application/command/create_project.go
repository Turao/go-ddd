package command

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type CreateProjectRequest struct {
	Title string
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
	p, err := project.NewProject(req.Title)
	if err != nil {
		return err
	}

	if err := h.repo.Create(ctx, *p); err != nil {
		return err
	}

	return nil
}
