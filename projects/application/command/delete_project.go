package command

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type DeleteProjectCommand struct {
	Title string `json:"title"`
}

type DeleteProjectHandler struct {
	repo project.Repository
}

func NewDeleteProjectCommandHandler(repo project.Repository) *DeleteProjectHandler {
	return &DeleteProjectHandler{
		repo: repo,
	}
}

func (h *DeleteProjectHandler) Handle(ctx context.Context, req DeleteProjectCommand) error {
	p, err := project.From(req.Title)
	if err != nil {
		return err
	}

	p.Delete()

	if err := h.repo.Save(ctx, *p); err != nil {
		return err
	}

	return nil
}
