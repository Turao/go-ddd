package in_memory

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/turao/go-ddd/events"
)

type MockEvent struct{}

var _ events.Event = (*MockEvent)(nil)

func (e MockEvent) ID() string           { return "" }
func (e MockEvent) Name() string         { return "" }
func (e MockEvent) OccuredAt() time.Time { return time.Now() }

func TestPush(t *testing.T) {
	type test struct {
		Event           events.Event
		ExpectedVersion int

		ExpectedSize int
		Error        error
	}

	tests := []test{
		{Event: &MockEvent{}, ExpectedVersion: 1, ExpectedSize: 1, Error: nil},
		{Event: &MockEvent{}, ExpectedVersion: 2, ExpectedSize: 0, Error: ErrExpectedVersionNotSatisfied},
	}

	for _, test := range tests {
		es, err := NewInMemoryStore()
		if err != nil {
			panic(err)
		}

		err = es.Push(context.Background(), test.Event, test.ExpectedVersion)
		assert.Equal(t, err, test.Error)
		if len(es.evts) != test.ExpectedSize {
			t.Errorf("event store should have %v event(s)", test.ExpectedSize)
		}
	}
}
