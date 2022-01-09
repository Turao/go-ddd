package command

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/projects/domain/project"
)

type UpdateProjectCommand struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type UpdateProjectHandler struct {
	repo project.Repository
}

func NewUpdateProjectCommandHandler(repo project.Repository) *UpdateProjectHandler {
	return &UpdateProjectHandler{
		repo: repo,
	}
}

func (h *UpdateProjectHandler) Handle(ctx context.Context, req UpdateProjectCommand) error {
	found, err := h.repo.FindProjectByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if found == nil {
		return errors.New("project does not exist")
	}

	updated, err := project.NewProject(
		found.ID,
		req.Title,
		found.Active,
	)

	if err != nil {
		return err
	}

	if err := h.repo.Save(ctx, *updated); err != nil {
		return err
	}

	return nil
}
