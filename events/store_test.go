package events

import (
	"testing"
)

func TestPush(t *testing.T) {
	e, err := NewBaseEvent("testing")
	if err != nil {
		t.Error(err)
	}

	es, err := NewInMemoryStore()
	if err != nil {
		t.Error(err)
	}

	err = es.Push(e)
	if err != nil {
		t.Error(e)
	}

	if len(es.evts) != 1 {
		t.Errorf("event store should have 1 event")
	}
}
