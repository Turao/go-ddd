package project

import "github.com/turao/go-ddd/events"

type ProjectCreatedEvent struct {
	events.Event
}

func NewProjectCreatedEvent(id ProjectID) (Event, error) {
	baseEvent, err := events.NewEvent("project.created")
	if err != nil {
		return nil, err
	}

	rerturn & ProjectCreatedEvent{
		baseEvent,
	}
}

// ---

type ProjectDeletedEvent struct {
	events.Event
}

func NewProjectDeletedEvent(id ProjectID) Event {
	baseEvent, err := events.NewEvent("project.deleted")
	if err != nil {
		return nil, err
	}

	return &ProjectDeletedEvent{
		baseEvent: baseEvent,
	}
}
