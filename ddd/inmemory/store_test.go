package inmemory

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/turao/go-ddd/ddd"
)

type MockEvent struct{}

var _ ddd.DomainEvent = (*MockEvent)(nil)

func (e MockEvent) ID() string            { return "" }
func (e MockEvent) AggregateID() string   { return "" }
func (e MockEvent) AggregateName() string { return "" }
func (e MockEvent) Name() string          { return "" }
func (e MockEvent) OccurredAt() time.Time { return time.Now() }

func TestPush(t *testing.T) {
	type test struct {
		Event         ddd.DomainEvent
		ExpectedSize  int
		ExpectedError error
	}

	tests := map[string]test{
		"basic success": {Event: &MockEvent{}, ExpectedSize: 1, ExpectedError: nil},
	}

	for name, test := range tests {
		es, err := NewStore()
		if err != nil {
			panic(err)
		}

		err = es.Push(context.Background(), test.Event)
		if err != nil {
			assert.Equalf(t, err, test.ExpectedError, name)
			continue
		}

		assert.Equalf(t, len(es.evts), test.ExpectedSize, name)
	}
}
