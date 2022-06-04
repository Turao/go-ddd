package user

import (
	"context"
	"errors"
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
		"unable to handle event": {
			Event:         func() ddd.DomainEvent { return nil },
			ExpectedError: errors.New("unable to handle event"),
		},
	}
	for name, test := range tests {
		agg := NewUserAggregate(UserEventsFactory{})
		err := agg.HandleEvent(context.Background(), test.Event())
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
		"unable to handle command": {
			Command:        nil,
			ExpectedEvents: nil,
			ExpectedError:  errors.New("unable to handle command"),
		},
	}
	for name, test := range tests {
		agg := NewUserAggregate(UserEventsFactory{})
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
			ExpectedError:  errors.New("invalid user name"),
		},
	}

	for name, test := range tests {
		// todo: mock event factory
		agg := NewUserAggregate(UserEventsFactory{})
		evts, err := agg.RegisterUser(test.Command)
		assert.Equalf(t, test.ExpectedEvents, evts, name)
		assert.Equalf(t, test.ExpectedError, err, name)
	}
}
