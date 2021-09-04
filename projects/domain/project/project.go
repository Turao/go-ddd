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

func NewProject(title string) (*Project, error) {
	if err := validateTitle(title); err != nil {
		return nil, err
	}

	return &Project{
		ID:    uuid.NewString(),
		Title: title,
	}, nil
}

func validateTitle(title string) error {
	if title == "" {
		return errors.New("title must not be empty")
	}
	return nil
}
