package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/projects/domain/project"
)

type ProjectRepository struct {
	projects map[project.ProjectID]*project.Project
}

var _ project.Repository = (*ProjectRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewProjectRepository() (*ProjectRepository, error) {
	return &ProjectRepository{}, nil
}

func (pr ProjectRepository) FindProjectByID(ctx context.Context, id project.ProjectID) (*project.Project, error) {
	p, found := pr.projects[id]
	if !found {
		return nil, ErrNotFound
	}

	return p, nil
}

func (pr ProjectRepository) Save(ctx context.Context, p project.Project) error {
	pr.projects[p.ID] = &p
	return nil
}
