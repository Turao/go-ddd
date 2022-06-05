package ddd

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAggregate struct {
	id string
}

var _ Aggregate = (*mockAggregate)(nil)

func (m mockAggregate) ID() string {
	return m.id
}
func (m mockAggregate) HandleEvent(ctx context.Context, evt DomainEvent) error {
	return nil
}
func (m mockAggregate) HandleCommand(ctx context.Context, cmd interface{}) ([]DomainEvent, error) {
	return nil, nil
}
func (m mockAggregate) MarshalJSON() ([]byte, error) {
	payload := struct {
		ID string `json:"id"`
	}{
		ID: m.id,
	}
	return json.Marshal(payload)
}
func (m *mockAggregate) UnmarshalJSON(data []byte) error {
	var payload struct {
		ID string `json:"id"`
	}
	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}
	m.id = payload.ID
	return nil
}

func TestSnapshot(t *testing.T) {
	// take a snapshot
	root, err := NewAggregateRoot(&mockAggregate{id: "mock-id"})
	if err != nil {
		t.Fatal(err)
	}
	snapshot, err := root.TakeSnapshot()
	assert.NoError(t, err)

	// restore from snapshot
	root2, err := NewAggregateRoot(&mockAggregate{})
	if err != nil {
		t.Fatal(err)
	}
	err = root2.FromSnapshot(snapshot)
	assert.NoError(t, err)

	assert.Equal(t, root.ID(), root2.ID())
	assert.Equal(t, root.Version(), root2.Version())
	assert.Equal(t, root.aggregate.ID(), root2.aggregate.ID())
}
