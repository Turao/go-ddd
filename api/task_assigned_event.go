package api

import (
	"context"

	"github.com/turao/go-ddd/events"
)

type TaskAssignedEventPublisher interface {
	Publish(ctx context.Context, event TaskAssignedEvent) error
}

type TaskAssignedEvent struct {
	IntegrationEvent

	TaskID string `json:"taskId"`
	UserID string `json:"userId"`
}

const (
	TaskAssignedEventName = "task.assigned"
)

// var (
// 	ErrInvalidTaskID = errors.New("invalid task id")
// 	ErrInvalidUserID = errors.New("invalid user id")
// )

func NewTaskAssignedEvent(correlationID string, taskID string, userID string) (*TaskAssignedEvent, error) {
	ie, err := events.NewIntegrationEvent(TaskAssignedEventName, taskID, correlationID)
	if err != nil {
		return nil, err
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	if userID == "" {
		return nil, ErrInvalidStatus
	}

	return &TaskAssignedEvent{
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
