package user

import "github.com/turao/go-ddd/events"

type UserRegisteredEvent struct {
	events.DomainEvent `json:"domainEvent"`

	Username string `json:"username"`
}

// var (
// 	ErrInvalidUserID   = errors.New("invalid user id")
// )

func NewUserRegisteredEvent(id string) (*UserRegisteredEvent, error) {
	domainEvent, err := events.NewDomainEvent("user.registered", id)
	if err != nil {
		return nil, err
	}

	return &UserRegisteredEvent{
		DomainEvent: domainEvent,
	}, nil
}
