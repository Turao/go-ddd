package project

import "github.com/turao/go-ddd/events"

type ProjectDeletedEvent struct {
	BaseEvent events.Event
}

func NewProjectDeletedEvent() Event {
	name := "project.deleted"

	return &ProjectDeletedEvent{
		BaseEvent: events.NewEvent(name),
	}
}
