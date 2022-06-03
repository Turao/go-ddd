package api

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/events"
)

type TaskStatusUpdatedEventPublisher interface {
	Publish(ctx context.Context, event TaskStatusUpdatedEvent) error
}

type TaskStatusUpdatedEvent struct {
	events.IntegrationEvent

	TaskID string `json:"taskId"`
	Status string `json:"status"`
}

var _ events.IntegrationEvent = (*TaskStatusUpdatedEvent)(nil)

func NewTaskStatusUpdatedEvent(correlationID string, taskID string, status string) (*TaskStatusUpdatedEvent, error) {
	event, err := events.NewEvent("task.status.updated")
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

	if status == "" {
		return nil, errors.New("invalid status")
	}

	return &TaskStatusUpdatedEvent{
		IntegrationEvent: ie,
		TaskID:           taskID,
		Status:           status,
	}, nil
}
