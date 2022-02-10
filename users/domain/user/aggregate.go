package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
)

type UserAggregate struct {
	User    *User
	version int
	events  events.EventStore
}

func NewUserAggregate(u *User, es events.EventStore) (*UserAggregate, error) {
	return &UserAggregate{
		User:    u,
		version: 0,
		events:  es,
	}, nil
}

func (ua *UserAggregate) HandleEvent(event events.DomainEvent) error {
	switch e := event.(type) {
	case UserRegisteredEvent:
		u, err := NewUser(e.AggregateID(), e.Username)
		if err != nil {
			return err
		}
		ua.User = u
		ua.version += 1
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}

}

func (ua *UserAggregate) RegisterUser(name string) error {
	u, err := NewUser(uuid.NewString(), name)
	if err != nil {
		return err
	}

	ua.User = u

	evt, err := NewUserRegisteredEvent(u.ID, u.Name)
	if err != nil {
		return err
	}

	err = ua.events.Push(context.Background(), *evt, ua.version+1)
	if err != nil {
		return err
	}
	ua.version += 1

	return nil
}
