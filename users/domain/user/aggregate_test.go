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

func (m *mockEventStore) Push(ctx context.Context, event events.Event) error {
	return nil
}

func (m *mockEventStore) Events(ctx context.Context) ([]events.Event, error) {
	return make([]events.Event, 0), nil
}

func (m *mockEventStore) EventsByAggregateID(ctx context.Context, aggregateID events.AggregateID) ([]events.DomainEvent, error) {
	return make([]events.DomainEvent, 0), nil
}

func TestRegisterUser(t *testing.T) {
	type test struct {
		inputName       string
		publishedEvents []events.Event

		fail bool
		err  error
	}

	tests := []test{
		{inputName: "dummy", publishedEvents: make([]events.Event, 0), fail: false, err: nil},
		{inputName: "", publishedEvents: make([]events.Event, 0), fail: true, err: ErrInvalidUserName},
	}

	for _, test := range tests {
		eventStore := new(mockEventStore)
		agg, err := NewUserAggregate(nil, eventStore)
		assert.NoError(t, err)

		err = agg.RegisterUser(test.inputName)

		if test.fail {
			assert.Equal(t, test.err, err)
			continue
		}

		assert.NoError(t, err)

		// todo: assert event is pushed to eventStore
	}
}
