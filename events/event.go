package events

import (
	"time"

	"github.com/google/uuid"
)

type EventID = string

type Event interface {
	ID() EventID
	Name() string
	OccuredAt() time.Time
}

type event struct {
	id        EventID
	name      string
	occuredAt time.Time
}

func NewEvent(name string) (*event, error) {
	return &event{
		id:        uuid.NewString(),
		name:      name,
		occuredAt: time.Now(),
	}, nil
}

func (e *event) ID() EventID {
	return e.id
}

func (e *event) Name() string {
	return e.name
}

func (e *event) OccuredAt() time.Time {
	return e.occuredAt
}
