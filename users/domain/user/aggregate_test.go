package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/turao/go-ddd/ddd"
)

func TestHandleEvent(t *testing.T) {
	tests := map[string]struct {
		Event         func() ddd.DomainEvent
		ExpectedError error
	}{
		"success": {
			Event: func() ddd.DomainEvent {
				evt, err := UserEventsFactory{}.NewUserRegisteredEvent("id", "name") // todo: don't use the real thing
				if err != nil {
					t.Fatal(err)
				}
				return evt
			},
			ExpectedError: nil,
		},
		"unknown  event": {
			Event:         func() ddd.DomainEvent { return nil },
			ExpectedError: ErrUnknownEvent,
		},
	}
	for name, test := range tests {
		agg, err := NewUserAggregate(UserEventsFactory{})
		if err != nil {
			t.Fatal(err)
		}

		err = agg.HandleEvent(context.Background(), test.Event())
		assert.Equalf(t, test.ExpectedError, err, name)
	}
}

func TestHandleCommand(t *testing.T) {
	tests := map[string]struct {
		Command        interface{}
		ExpectedEvents []ddd.DomainEvent
		ExpectedError  error
	}{
		"success": {
			Command:        RegisterUserCommand{Username: "dummy"},
			ExpectedEvents: []ddd.DomainEvent{},
			ExpectedError:  nil,
		},
		"unknown command": {
			Command:        nil,
			ExpectedEvents: nil,
			ExpectedError:  ErrUnknownCommand,
		},
	}
	for name, test := range tests {
		agg, err := NewUserAggregate(UserEventsFactory{})
		if err != nil {
			t.Fatal(err)
		}

		evts, err := agg.HandleCommand(context.Background(), test.Command)
		assert.Equalf(t, test.ExpectedEvents, evts, name)
		assert.Equalf(t, test.ExpectedError, err, name)
	}
}

func TestRegisterUser(t *testing.T) {
	type test struct {
		Command        RegisterUserCommand
		ExpectedEvents []ddd.DomainEvent
		ExpectedError  error
	}

	tests := map[string]test{
		"success": {
			Command:        RegisterUserCommand{Username: "dummy"},
			ExpectedEvents: []ddd.DomainEvent{},
			ExpectedError:  nil,
		},
		"empty user name": {
			Command:        RegisterUserCommand{Username: ""},
			ExpectedEvents: []ddd.DomainEvent{},
			ExpectedError:  ErrEmptyUserName,
		},
	}

	for name, test := range tests {
		// todo: mock event factory
		agg, err := NewUserAggregate(UserEventsFactory{})
		if err != nil {
			t.Fatal(err)
		}

		evts, err := agg.handleRegisterUserCommand(test.Command)
		assert.Equalf(t, test.ExpectedEvents, evts, name)
		assert.Equalf(t, test.ExpectedError, err, name)
	}
}
