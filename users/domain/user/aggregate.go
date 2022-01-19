package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
)

type UserAggregate struct {
	User *User

	events events.EventStore
}

func NewUserAggregate(u *User, es events.EventStore) (*UserAggregate, error) {
	return &UserAggregate{
		User:   u,
		events: es,
	}, nil
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

	err = ua.events.Push(context.Background(), *evt)
	if err != nil {
		return err
	}

	return nil
}
