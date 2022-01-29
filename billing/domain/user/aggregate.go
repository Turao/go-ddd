package user

import (
	"context"
	"fmt"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/users/domain/user"
)

type UserAggregate struct {
	User   *User
	events events.EventStore
}

func NewUserAggregate(user *User, es events.EventStore) (*UserAggregate, error) {
	return &UserAggregate{
		User:   user,
		events: es,
	}, nil
}

func (ua *UserAggregate) HandleEvent(event events.DomainEvent) error {
	switch e := event.(type) {
	case UserRegisteredEvent:
		u, err := NewUser(e.AggregateID())
		if err != nil {
			return err
		}
		ua.User = u
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}
}

func (ua *UserAggregate) RegisterUser(userID user.UserID) error {
	u, err := NewUser(userID)
	if err != nil {
		return err
	}

	ua.User = u

	evt, err := NewUserRegisteredEvent(u.ID)
	if err != nil {
		return err
	}

	err = ua.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}
