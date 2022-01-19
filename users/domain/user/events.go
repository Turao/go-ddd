package user

import "github.com/turao/go-ddd/events"

type UserRegisteredEvent struct {
	events.DomainEvent `json:"domainEvent"`

	UserName string `json:"userName"`
}

// var (
// 	ErrInvalidUserID   = errors.New("invalid user id")
// 	ErrInvalidUserName = errors.New("invalid user name")
// )

func NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error) {
	domainEvent, err := events.NewDomainEvent("user.registered", id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, ErrInvalidUserName
	}

	return &UserRegisteredEvent{
		DomainEvent: domainEvent,
		UserName:    name,
	}, nil
}
