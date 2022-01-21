package in_memory

import (
	"context"
	"testing"
	"time"

	"github.com/turao/go-ddd/events"
)

type MockEvent struct{}

var _ events.Event = (*MockEvent)(nil)

func (e MockEvent) ID() string           { return "" }
func (e MockEvent) Name() string         { return "" }
func (e MockEvent) OccuredAt() time.Time { return time.Now() }

func TestPush(t *testing.T) {
	e := &MockEvent{}

	es, err := NewInMemoryStore()
	if err != nil {
		t.Error(err)
	}

	err = es.Push(context.Background(), e)
	if err != nil {
		t.Error(e)
	}

	if len(es.evts) != 1 {
		t.Errorf("event store should have 1 event")
	}
}
