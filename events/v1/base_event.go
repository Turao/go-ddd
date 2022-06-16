package v1

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
)

type BaseEvent struct {
	id        string
	name      string
	occuredAt time.Time
}

var _ events.Event = (*BaseEvent)(nil)

var (
	ErrInvalidName = errors.New("invalid event name")
)

func NewEvent(name string) (*BaseEvent, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	return &BaseEvent{
		id:        uuid.NewString(),
		name:      name,
		occuredAt: time.Now(),
	}, nil
}

func (e BaseEvent) ID() string {
	return e.id
}

func (e BaseEvent) Name() string {
	return e.name
}

func (e BaseEvent) OccurredAt() time.Time {
	return e.occuredAt
}

func (e BaseEvent) MarshalJSON() ([]byte, error) {
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

func (e *BaseEvent) UnmarshalJson(data []byte) error {
	payload := struct {
		ID         string    `json:"id"`
		Name       string    `json:"name"`
		OccurredAt time.Time `json:"occurredAt"`
	}{}

	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}

	e.id = payload.ID
	e.name = payload.Name
	e.occuredAt = payload.OccurredAt
	return nil
}
