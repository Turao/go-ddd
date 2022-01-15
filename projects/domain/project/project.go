package project

import (
	"errors"
)

type ProjectID = string

var (
	ErrInvalidProjectID   = errors.New("invalid project id")
	ErrInvalidProjectName = errors.New("invalid project name")
)

type Project struct {
	ID   ProjectID `json:"id"`
	Name string    `json:"name"`

	Active bool `json:"active"`
}

func NewProject(id ProjectID, name string, active bool) (*Project, error) {
	if id == "" {
		return nil, ErrInvalidProjectID
	}

	if err := validateName(name); err != nil {
		return nil, err
	}

	return &Project{
		ID:     id,
		Name:   name,
		Active: active,
	}, nil
}

func validateName(name string) error {
	if name == "" {
		return ErrInvalidProjectName
	}
	return nil
}

func (p *Project) Rename(name string) error {
	if err := validateName(name); err != nil {
		return err
	}

	p.Name = name
	return nil
}

func (p *Project) Delete() {
	p.Active = false
}
