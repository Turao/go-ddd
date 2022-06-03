package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/turao/go-ddd/ddd"
)

type mockEventFactory struct{}

func (ef *mockEventFactory) NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error) {
	return &UserRegisteredEvent{}, nil
}

func TestRegisterUser(t *testing.T) {
	type test struct {
		InputName      string
		ExpectedEvents func() []ddd.DomainEvent
		ExpectedError  error
	}

	tests := map[string]test{
		"success": {
			InputName: "dummy",
			ExpectedEvents: func() []ddd.DomainEvent {
				return make([]ddd.DomainEvent, 0)
			},
			ExpectedError: nil,
		},
		"empty user name": {
			InputName: "",
			ExpectedEvents: func() []ddd.DomainEvent {
				return make([]ddd.DomainEvent, 0)
			},
			ExpectedError: errors.New("invalid user name"),
		},
	}

	for _, test := range tests {
		agg := NewUserAggregate(&mockEventFactory{})
		evts, err := agg.RegisterUser(RegisterUserCommand{test.InputName})

		assert.EqualValues(t, test.ExpectedEvents, evts) // todo: fix mocking of event factory
		assert.Equal(t, test.ExpectedError, err)
	}
}
