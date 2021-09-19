package command

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type DeleteProjectRequest struct {
	Title string `json:"title"`
}

type DeleteProjectHandler struct {
	repo project.WriteRepository
}

func NewDeleteProjectCommandHandler(repo project.WriteRepository) *DeleteProjectHandler {
	return &DeleteProjectHandler{
		repo: repo,
	}
}

func (h *DeleteProjectHandler) Handle(ctx context.Context, req DeleteProjectRequest) error {
	p, err := project.From(req.Title)
	if err != nil {
		return err
	}

	if err := h.repo.Delete(ctx, p.ID); err != nil {
		return err
	}

	return nil
}
