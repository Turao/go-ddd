package api

import (
	"context"

	"github.com/turao/go-ddd/events"
)

type TaskUnassignedEventPublisher interface {
	Publish(ctx context.Context, event TaskUnassignedEvent) error
}

type TaskUnassignedEvent struct {
	events.IntegrationEvent

	TaskID string `json:"taskId"`
	UserID string `json:"userId"`
}

var _ events.IntegrationEvent = (*TaskUnassignedEvent)(nil)

const (
	TaskUnassignedEventName = "task.unassigned"
)

// var (
// 	ErrInvalidTaskID = errors.New("invalid task id")
// 	ErrInvalidUserID = errors.New("invalid user id")
// )

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
		return nil, ErrInvalidTaskID
	}

	if userID == "" {
		return nil, ErrInvalidStatus
	}

	return &TaskUnassignedEvent{
		IntegrationEvent: ie,
		TaskID:           taskID,
		UserID:           userID,
	}, nil
}
