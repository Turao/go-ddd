package project

import (
	"errors"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/domain/user"
)

type ProjectCreatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	ProjectName        string      `json:"projectName"`
	CreatedBy          user.UserID `json:"createdBy"`
}

func NewProjectCreatedEvent(id ProjectID, projectName string, createdBy user.UserID) (*ProjectCreatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.created", id)
	if err != nil {
		return nil, err
	}

	if projectName == "" {
		return nil, errors.New("project name must not be empty")
	}

	if createdBy == "" {
		return nil, ErrInvalidUserID
	}

	return &ProjectCreatedEvent{
		DomainEvent: domainEvent,
		ProjectName: projectName,
		CreatedBy:   createdBy,
	}, nil
}

type ProjectUpdatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	ProjectName        string `json:"projectName"`
}

func NewProjectUpdatedEvent(id ProjectID, projectName string) (*ProjectUpdatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.updated", id)
	if err != nil {
		return nil, err
	}

	if projectName == "" {
		return nil, errors.New("project name must not be empty")
	}

	return &ProjectUpdatedEvent{
		domainEvent,
		projectName,
	}, nil
}

type ProjectDeletedEvent struct {
	events.DomainEvent `json:"domainEvent"`
}

func NewProjectDeletedEvent(id ProjectID) (*ProjectDeletedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.deleted", id)
	if err != nil {
		return nil, err
	}

	return &ProjectDeletedEvent{
		domainEvent,
	}, nil
}
