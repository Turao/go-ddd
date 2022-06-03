package api

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/events"
)

type TaskAssignedEventPublisher interface {
	Publish(ctx context.Context, event TaskAssignedEvent) error
}

type TaskAssignedEvent struct {
	events.IntegrationEvent

	TaskID string `json:"taskId"`
	UserID string `json:"userId"`
}

var _ events.IntegrationEvent = (*TaskAssignedEvent)(nil)

func NewTaskAssignedEvent(correlationID string, taskID string, userID string) (*TaskAssignedEvent, error) {
	event, err := events.NewEvent("task.assigned")
	if err != nil {
		return nil, err
	}

	ie, err := events.NewIntegrationEvent(event, correlationID)
	if err != nil {
		return nil, err
	}

	if taskID == "" {
		return nil, errors.New("invalid task id")
	}

	if userID == "" {
		return nil, errors.New("invalid status")
	}

	return &TaskAssignedEvent{
		IntegrationEvent: ie,
		TaskID:           taskID,
		UserID:           userID,
	}, nil
}
