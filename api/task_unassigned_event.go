package api

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/events"
	v1 "github.com/turao/go-ddd/events/v1"
)

type TaskUnassignedEventPublisher interface {
	Publish(ctx context.Context, event TaskUnassignedEvent) error
}

type TaskUnassignedEvent struct {
	*v1.IntegrationEvent

	TaskID string `json:"taskId"`
	UserID string `json:"userId"`
}

var _ events.IntegrationEvent = (*TaskUnassignedEvent)(nil)

const TaskUnassignedEventName = "task.unassigned"

func NewTaskUnassignedEvent(correlationID string, taskID string, userID string) (*TaskUnassignedEvent, error) {
	event, err := v1.NewEvent(TaskUnassignedEventName)
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

	if userID == "" {
		return nil, errors.New("invalid status")
	}

	return &TaskUnassignedEvent{
		IntegrationEvent: ie,
		TaskID:           taskID,
		UserID:           userID,
	}, nil
}
