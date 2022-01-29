package domain

import "github.com/turao/go-ddd/events"

type AggregateRoot struct {
	User *User
}

func NewAggregateRoot(user *User) (*AggregateRoot, error) {
	return &AggregateRoot{
		User: user,
	}, nil
}

func (a AggregateRoot) HandleEvent(event events.DomainEvent) error {
	return nil
}
