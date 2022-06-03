package events

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type event struct {
	id        string
	name      string
	occuredAt time.Time
}

var _ Event = (*event)(nil)

var (
	ErrInvalidName = errors.New("invalid event name")
)

func NewEvent(name string) (*event, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	return &event{
		id:        uuid.NewString(),
		name:      name,
		occuredAt: time.Now(),
	}, nil
}

func (e event) ID() string {
	return e.id
}

func (e event) Name() string {
	return e.name
}

func (e event) OccurredAt() time.Time {
	return e.occuredAt
}

func (e event) MarshalJSON() ([]byte, error) {
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
