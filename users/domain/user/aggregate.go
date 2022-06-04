package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
)

type UserAggregate struct {
	User *User

	EventFactory
}

var (
	ErrUnknownEvent   = errors.New("unknown event")
	ErrUnknownCommand = errors.New("unknown command")
)

func NewUserAggregate(ef EventFactory) *UserAggregate {
	return &UserAggregate{
		User:         nil,
		EventFactory: ef,
	}
}

func (ua UserAggregate) ID() string {
	return ua.User.ID
}

func (ua *UserAggregate) HandleEvent(ctx context.Context, event ddd.DomainEvent) error {
	switch e := event.(type) {
	case UserRegisteredEvent:
		u, err := NewUser(e.AggregateID(), e.Username)
		if err != nil {
			return err
		}
		ua.User = u
		return nil
	default:
		return ErrUnknownEvent
	}

}

func (ua *UserAggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]ddd.DomainEvent, error) {
	switch c := cmd.(type) {
	case RegisterUserCommand:
		return ua.RegisterUser(c)
	default:
		return nil, ErrUnknownCommand
	}
}

func (ua *UserAggregate) RegisterUser(cmd RegisterUserCommand) ([]ddd.DomainEvent, error) {
	u, err := NewUser(uuid.NewString(), cmd.Username)
	if err != nil {
		return nil, err
	}

	ua.User = u

	evt, err := ua.EventFactory.NewUserRegisteredEvent(u.ID, u.Name)
	if err != nil {
		return nil, err
	}

	return []ddd.DomainEvent{
		*evt,
	}, nil
}
