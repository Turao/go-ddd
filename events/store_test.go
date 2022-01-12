package events

import (
	"context"
	"testing"
)

func TestPush(t *testing.T) {
	e, err := NewDomainEvent("testing", "aggregate-id")
	if err != nil {
		t.Error(err)
	}

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
