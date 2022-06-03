package user

import (
	"github.com/turao/go-ddd/ddd"
	"github.com/turao/go-ddd/events"
)

type EventFactory interface {
	NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error)
}

type UserEventsFactory struct{}

type UserRegisteredEvent struct {
	ddd.DomainEvent `json:"domainEvent"`

	Username string `json:"username"`
}

// var (
// 	ErrInvalidUserID   = errors.New("invalid user id")
// 	ErrInvalidUserName = errors.New("invalid user name")
// )

func (f UserEventsFactory) NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error) {
	event, err := events.NewEvent("user.registered")
	if err != nil {
		return nil, err
	}

	domainEvent, err := ddd.NewDomainEvent(event, id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, ErrInvalidUserName
	}

	return &UserRegisteredEvent{
		DomainEvent: domainEvent,
		Username:    name,
	}, nil
}
