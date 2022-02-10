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

	tests := map[string]test{
		"basic success":          {Event: &MockEvent{}, ExpectedVersion: 1, ExpectedSize: 1, Error: nil},
		"wrong expected version": {Event: &MockEvent{}, ExpectedVersion: 2, ExpectedSize: 0, Error: ErrExpectedVersionNotSatisfied},
	}

	for name, test := range tests {
		es, err := NewInMemoryStore()
		if err != nil {
			panic(err)
		}

		err = es.Push(context.Background(), test.Event, test.ExpectedVersion)
		if err != nil {
			assert.Equalf(t, err, test.Error, name)
			continue
		}

		assert.Equalf(t, len(es.evts), test.ExpectedSize, name)
	}
}
