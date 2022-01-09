package events

import (
	"time"
)

type EventID = string

type Event interface {
	ID() EventID
	Name() string
	OccuredAt() time.Time
}
