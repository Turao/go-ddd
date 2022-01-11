package project

import (
	"errors"

	"github.com/google/uuid"
)

type ProjectID = string

type Project struct {
	ID   ProjectID `json:"id"`
	Name string    `json:"name"`

	Active bool `json:"active"`
}

func From(name string) (*Project, error) {
	return NewProject(uuid.NewString(), name, true)
}

func NewProject(id ProjectID, name string, active bool) (*Project, error) {
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
		return errors.New("name must not be empty")
	}
	return nil
}

func (p *Project) SetName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}

	p.Name = name
	return nil
}

func (p *Project) Delete() {
	p.Active = false
}
