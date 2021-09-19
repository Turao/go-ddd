package project

import (
	"errors"

	"github.com/google/uuid"
)

type ProjectID = string

type Project struct {
	ID    ProjectID
	Title string
}

func NewProject(id ProjectID, title string) (*Project, error) {
	if err := validateTitle(title); err != nil {
		return nil, err
	}

	return &Project{
		ID:    id,
		Title: title,
	}, nil
}

func validateTitle(title string) error {
	if title == "" {
		return errors.New("title must not be empty")
	}
	return nil
}

func From(title string) (*Project, error) {
	return NewProject(uuid.NewString(), title)
}
