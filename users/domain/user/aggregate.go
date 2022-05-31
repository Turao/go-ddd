package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
)

type UserAggregate struct {
	User *User

	EventFactory
}

func NewUserAggregate(ef EventFactory) *UserAggregate {
	return &UserAggregate{
		User:         nil,
		EventFactory: ef,
	}
}

func (ua UserAggregate) ID() string {
	return ua.User.ID
}

func (ua *UserAggregate) HandleEvent(ctx context.Context, event events.DomainEvent) error {
	switch e := event.(type) {
	case UserRegisteredEvent:
		u, err := NewUser(e.AggregateID(), e.Username)
		if err != nil {
			return err
		}
		ua.User = u
		return nil
	default:
		return fmt.Errorf("unable to handle domain event %s", e)
	}

}

func (ua *UserAggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]events.DomainEvent, error) {
	switch c := cmd.(type) {
	case RegisterUserCommand:
		return ua.RegisterUser(c)
	default:
		return nil, fmt.Errorf("unable to handle command %s", cmd)
	}
}

func (ua *UserAggregate) RegisterUser(cmd RegisterUserCommand) ([]events.DomainEvent, error) {
	u, err := NewUser(uuid.NewString(), cmd.Username)
	if err != nil {
		return nil, err
	}

	ua.User = u

	evt, err := ua.EventFactory.NewUserRegisteredEvent(u.ID, u.Name)
	if err != nil {
		return nil, err
	}

	return []events.DomainEvent{
		*evt,
	}, nil
}
