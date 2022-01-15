package project

import (
	"errors"

	"github.com/turao/go-ddd/events"
)

type ProjectCreatedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	ProjectName        string `json:"projectName"`
}

func NewProjectCreatedEvent(id ProjectID, projectName string) (*ProjectCreatedEvent, error) {
	domainEvent, err := events.NewDomainEvent("project.created", id)
	if err != nil {
		return nil, err
	}

	if projectName == "" {
		return nil, errors.New("project name must not be empty")
	}

	return &ProjectCreatedEvent{
		domainEvent,
		projectName,
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
