package api

import (
	"context"

	"github.com/turao/go-ddd/events"
)

type TaskUnassignedEventPublisher interface {
	Publish(ctx context.Context, event TaskUnassignedEvent) error
}

type TaskUnassignedEvent struct {
	IntegrationEvent

	TaskID string `json:"taskId"`
	UserID string `json:"userId"`
}

const (
	TaskUnassignedEventName = "task.unassigned"
)

// var (
// 	ErrInvalidTaskID = errors.New("invalid task id")
// 	ErrInvalidUserID = errors.New("invalid user id")
// )

func NewTaskUnassignedEvent(correlationID string, taskID string, userID string) (*TaskUnassignedEvent, error) {
	ie, err := events.NewIntegrationEvent(TaskUnassignedEventName, taskID, correlationID)
	if err != nil {
		return nil, err
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	if userID == "" {
		return nil, ErrInvalidStatus
	}

	return &TaskUnassignedEvent{
		IntegrationEvent: IntegrationEvent{
			DomainEvent: DomainEvent{
				BaseEvent: BaseEvent{
					ID:         ie.ID(),
					Name:       ie.Name(),
					OccurredAt: ie.OccuredAt(),
				},
				AggregateID: ie.AggregateID(),
			},
			CorrelationID: ie.CorrelationID(),
		},

		TaskID: taskID,
		UserID: userID,
	}, nil
}
