package events

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type baseEvent struct {
	id        EventID
	name      string
	occuredAt time.Time
}

func NewBaseEvent(name string) (*baseEvent, error) {
	if name == "" {
		return nil, errors.New("event name must not be empty")
	}

	return &baseEvent{
		id:        uuid.NewString(),
		name:      name,
		occuredAt: time.Now(),
	}, nil
}

func (e *baseEvent) ID() EventID {
	return e.id
}

func (e *baseEvent) Name() string {
	return e.name
}

func (e *baseEvent) OccuredAt() time.Time {
	return e.occuredAt
}
