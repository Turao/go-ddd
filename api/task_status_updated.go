package api

import (
	"errors"

	"github.com/turao/go-ddd/events"
)

type TaskStatusUpdatedEvent struct {
	IntegrationEvent

	TaskID string `json:"taskId"`
	Status string `json:"status"`
}

const (
	TaskStatusUpdatedEventName = "task.status.updated"
)

var (
	ErrInvalidTaskID = errors.New("invalid task id")
	ErrInvalidStatus = errors.New("invalid status")
)

func NewTaskStatusUpdatedEvent(correlationID string, taskID string, status string) (*TaskStatusUpdatedEvent, error) {
	ie, err := events.NewIntegrationEvent(TaskStatusUpdatedEventName, taskID, correlationID)
	if err != nil {
		return nil, err
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	if status == "" {
		return nil, ErrInvalidStatus
	}

	return &TaskStatusUpdatedEvent{
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
		Status: status,
	}, nil
}
