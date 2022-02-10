package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/turao/go-ddd/events"
)

type mockEventStore struct {
	mock.Mock
}

var _ events.EventStore = (*mockEventStore)(nil)

func (m *mockEventStore) Push(ctx context.Context, event events.Event, expectedVersion int) error {
	args := m.Called(ctx, event, expectedVersion)
	return args.Error(0)
}

func (m *mockEventStore) Events(ctx context.Context) ([]events.Event, error) {
	return make([]events.Event, 0), nil
}

func (m *mockEventStore) EventsByAggregateID(ctx context.Context, aggregateID events.AggregateID) ([]events.DomainEvent, error) {
	return make([]events.DomainEvent, 0), nil
}

func TestRegisterUser(t *testing.T) {
	type test struct {
		InputName       string
		PublishedEvents []events.Event

		Error error
	}

	tests := map[string]test{
		"success":         {InputName: "dummy", PublishedEvents: make([]events.Event, 0), Error: nil},
		"empty user name": {InputName: "", PublishedEvents: make([]events.Event, 0), Error: ErrInvalidUserName},
	}

	for name, test := range tests {
		eventStore := new(mockEventStore)
		agg, err := NewUserAggregate(nil, eventStore)
		assert.NoError(t, err)

		eventStore.On("Push", mock.Anything, mock.Anything, mock.Anything).Return(test.Error)

		err = agg.RegisterUser(test.InputName)

		if err != nil {
			assert.Equalf(t, err, test.Error, name)
			continue
		}

		// todo: should we assert on the event type?
		eventStore.AssertCalled(t, "Push", mock.Anything, mock.Anything, mock.Anything)
	}
}
