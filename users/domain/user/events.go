package user

import (
	"errors"

	"github.com/turao/go-ddd/ddd"
	v1 "github.com/turao/go-ddd/events/v1"
)

type EventFactory interface {
	NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error)
}

type UserEventsFactory struct{}

type UserRegisteredEvent struct {
	ddd.DomainEvent `json:"domainEvent"`

	Username string `json:"username"`
}

func (f UserEventsFactory) NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error) {
	event, err := v1.NewEvent("user.registered")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, id, UserAggregateName)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, errors.New("invalid user name")
	}

	return &UserRegisteredEvent{
		DomainEvent: domainEvent,
		Username:    name,
	}, nil
}
