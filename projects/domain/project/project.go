package project

import (
	"errors"
	"time"

	"github.com/turao/go-ddd/users/domain/user"
)

type ProjectID = string

var (
	ErrInvalidProjectID   = errors.New("invalid project id")
	ErrInvalidProjectName = errors.New("invalid project name")
	ErrInvalidUserID      = errors.New("invalid user id")
)

type Project struct {
	ID   ProjectID `json:"id"`
	Name string    `json:"name"`

	CreatedBy user.UserID `json:"createdBy"`
	CreatedAt time.Time   `json:"createdAt"`

	Active bool `json:"active"`
}

func NewProject(id ProjectID, name string, createdBy user.UserID, createdAt time.Time, active bool) (*Project, error) {
	if id == "" {
		return nil, ErrInvalidProjectID
	}

	if err := validateName(name); err != nil {
		return nil, err
	}

	if createdBy == "" {
		return nil, ErrInvalidUserID
	}

	return &Project{
		ID:        id,
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: createdAt,
		Active:    active,
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

func (p *Project) Delete() error {
	p.Active = false
	return nil
}
