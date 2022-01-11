package events

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type baseEvent struct {
	id        string
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

func (e *baseEvent) ID() string {
	return e.id
}

func (e *baseEvent) Name() string {
	return e.name
}

func (e *baseEvent) OccuredAt() time.Time {
	return e.occuredAt
}

func (e *baseEvent) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		OccuredAt time.Time `json:"occurredAt"`
	}{
		ID:        e.id,
		Name:      e.name,
		OccuredAt: e.occuredAt,
	})
	if err != nil {
		return nil, err
	}
	return d, err
}
