package ddd

import "github.com/turao/go-ddd/events"

type DomainEvent interface {
	events.Event
	AggregateID() string
}
