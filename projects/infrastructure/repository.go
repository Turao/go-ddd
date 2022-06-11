package infrastructure

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/projects/domain/project"
)

type ProjectRepository struct {
	projects map[project.ProjectID]*project.ProjectAggregate
}

var _ project.Repository = (*ProjectRepository)(nil)

var (
	ErrNotFound = errors.New("not found")
)

func NewProjectRepository() (*ProjectRepository, error) {
	return &ProjectRepository{
		projects: make(map[string]*project.ProjectAggregate),
	}, nil
}

func (pr ProjectRepository) FindProjectByID(ctx context.Context, id project.ProjectID) (*project.ProjectAggregate, error) {
	p, found := pr.projects[id]
	if !found {
		return nil, ErrNotFound
	}

	return p, nil
}

func (pr ProjectRepository) Save(ctx context.Context, p *project.ProjectAggregate) error {
	pr.projects[p.ID()] = p
	return nil
}

func (pr ProjectRepository) FindAll(ctx context.Context) ([]*project.ProjectAggregate, error) {
	var ps []*project.ProjectAggregate
	for _, p := range pr.projects {
		ps = append(ps, p)
	}
	return ps, nil
}
