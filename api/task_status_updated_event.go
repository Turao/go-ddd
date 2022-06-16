package api

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/events"
	v1 "github.com/turao/go-ddd/events/v1"
)

type TaskStatusUpdatedEventPublisher interface {
	Publish(ctx context.Context, event TaskStatusUpdatedEvent) error
}

type TaskStatusUpdatedEvent struct {
	*v1.IntegrationEvent

	TaskID string `json:"taskId"`
	Status string `json:"status"`
}

var _ events.IntegrationEvent = (*TaskStatusUpdatedEvent)(nil)

const TaskStatusUpdatedEventName = "task.status.updated"

func NewTaskStatusUpdatedEvent(correlationID string, taskID string, status string) (*TaskStatusUpdatedEvent, error) {
	event, err := v1.NewEvent(TaskStatusUpdatedEventName)
	if err != nil {
		return nil, err
	}

	ie, err := v1.NewIntegrationEvent(event, correlationID)
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
