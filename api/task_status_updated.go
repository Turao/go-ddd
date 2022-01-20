package api

import (
	"errors"

	"github.com/google/uuid"
	"github.com/turao/go-ddd/events"
)

type TaskStatusUpdatedEvent struct {
	events.IntegrationEvent
	TaskID string `json:"taskId"`
	Status string `json:"status"`
}

var (
	ErrInvalidTaskID = errors.New("invalid task id")
	ErrInvalidStatus = errors.New("invalid status")
)

func NewTaskStatusUpdatedEvent(taskID string, status string) (*TaskStatusUpdatedEvent, error) {
	integrationEvent, err := events.NewIntegrationEvent("task.status.updated", uuid.NewString())
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
		IntegrationEvent: integrationEvent,
		TaskID:           taskID,
		Status:           status,
	}, nil
}
