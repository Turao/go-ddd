package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/turao/go-ddd/ddd"
)

type mockEventFactory struct{}

func (ef *mockEventFactory) NewUserRegisteredEvent(id string, name string) (*UserRegisteredEvent, error) {
	return nil, nil
}

func TestRegisterUser(t *testing.T) {
	type test struct {
		InputName       string
		PublishedEvents []ddd.DomainEvent

		Error error
	}

	tests := map[string]test{
		"success":         {InputName: "dummy", PublishedEvents: nil, Error: nil}, // todo: fix this test
		"empty user name": {InputName: "", PublishedEvents: nil, Error: ErrInvalidUserName},
	}

	for _, test := range tests {
		agg := NewUserAggregate(&mockEventFactory{})
		_, err := agg.RegisterUser(RegisterUserCommand{test.InputName})
		assert.Equal(t, err, test.Error)
	}
}
