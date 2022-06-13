package api

import (
	"context"
	"errors"

	"github.com/turao/go-ddd/events"
)

type TaskUnassignedEventPublisher interface {
	Publish(ctx context.Context, event TaskUnassignedEvent) error
}

type TaskUnassignedEvent struct {
	*events.IntegrationEvent

	TaskID string `json:"taskId"`
	UserID string `json:"userId"`
}

// var _ events.IntegrationEvent = (*TaskUnassignedEvent)(nil)

const TaskUnassignedEventName = "task.unassigned"

func NewTaskUnassignedEvent(correlationID string, taskID string, userID string) (*TaskUnassignedEvent, error) {
	event, err := events.NewEvent(TaskUnassignedEventName)
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

	return &TaskUnassignedEvent{
		IntegrationEvent: ie,
		TaskID:           taskID,
		UserID:           userID,
	}, nil
}
