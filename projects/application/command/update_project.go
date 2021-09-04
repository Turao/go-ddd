package command

import (
	"context"

	"github.com/turao/go-ddd/projects/domain/project"
)

type UpdateProjectRequest struct {
	ID    string
	Title string
}

type UpdateProjectHandler struct {
	repo project.WriteRepository
}

func NewUpdateProjectCommandHandler(repo project.WriteRepository) *UpdateProjectHandler {
	return &UpdateProjectHandler{
		repo: repo,
	}
}

func (h *UpdateProjectHandler) Handle(ctx context.Context, req UpdateProjectRequest) error {
	p := project.Project{
		ID:    req.ID,
		Title: req.Title,
	}

	if err := h.repo.Update(ctx, p); err != nil {
		return err
	}

	return nil
}
