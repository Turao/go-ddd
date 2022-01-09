package events

import (
	"context"
	"testing"
)

func TestPush(t *testing.T) {
	e, err := NewEvent("testing")
	if err != nil {
		t.Error(err)
	}

	ims, err := NewInMemoryStore()
	if err != nil {
		t.Error(err)
	}

	err = ims.Push(context.Background(), e)
	if err != nil {
		t.Error(e)
	}

	if len(ims.events) != 1 {
		t.Errorf("event store should have 1 event")
	}
}

func TestTake(t *testing.T) {
	e, err := NewEvent("testing")
	if err != nil {
		t.Error(err)
	}

	ims, err := NewInMemoryStore()
	if err != nil {
		t.Error(err)
	}

	ims.events = []Event{e}

	taken := ims.Take(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}

	if len(taken) != 1 {
		t.Errorf("event store should have returned 1 event")
	}

	if len(ims.events) != 0 {
		t.Errorf("event store should have 1 event")
	}
}

func TestTakeMoreThanPossible(t *testing.T) {
	e, err := NewEvent("testing")
	if err != nil {
		t.Error(err)
	}

	ims, err := NewInMemoryStore()
	if err != nil {
		t.Error(err)
	}

	ims.events = []Event{e}

	taken := ims.Take(context.Background(), 2)
	if err != nil {
		t.Error(err)
	}

	if len(taken) != 1 {
		t.Errorf("event store should have returned 1 event")
	}

	if len(ims.events) != 0 {
		t.Errorf("event store should have 1 event")
	}
}
