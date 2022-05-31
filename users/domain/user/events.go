package user

import "github.com/turao/go-ddd/events"

type EventFactory interface {
	NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error)
}

type UserEventsFactory struct{}

type UserRegisteredEvent struct {
	events.DomainEvent `json:"domainEvent"`

	Username string `json:"username"`
}

// var (
// 	ErrInvalidUserID   = errors.New("invalid user id")
// 	ErrInvalidUserName = errors.New("invalid user name")
// )

func (f UserEventsFactory) NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error) {
	domainEvent, err := events.NewDomainEvent("user.registered", id)
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
