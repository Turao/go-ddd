package inmemory

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ThreeDotsLabs/watermill"
	watermillKafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/turao/go-ddd/ddd"
)

type store struct {
	publisher *watermillKafka.Publisher
}

var _ ddd.DomainEventStore = (*store)(nil)

var (
	ErrExpectedVersionNotSatisfied = errors.New("expected version does not match event store state")
)

func NewStore() (*store, error) {
	queue := watermillKafka.PublisherConfig{
		Brokers:   []string{"localhost:29092"},
		Marshaler: watermillKafka.DefaultMarshaler{},
	}
	logger := watermill.NewStdLogger(false, false)
	publisher, err := watermillKafka.NewPublisher(queue, logger)
	if err != nil {
		return nil, err
	}

	return &store{
		publisher: publisher,
	}, nil
}

func (s store) Push(ctx context.Context, evt ddd.DomainEvent) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return s.publisher.Publish(
		evt.AggregateName(),
		message.NewMessage(uuid.NewString(), data),
	)
}

func (s store) Events(ctx context.Context, aggregateID string) ([]ddd.DomainEvent, error) {
	return make([]ddd.DomainEvent, 0), nil
}

func (s store) Close() error {
	return s.publisher.Close()
}
