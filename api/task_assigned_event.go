package api

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/turao/go-ddd/events"
	v1 "github.com/turao/go-ddd/events/v1"
)

type TaskAssignedEventPublisher interface {
	Publish(ctx context.Context, event TaskAssignedEvent) error
}

type TaskAssignedEvent struct {
	*v1.IntegrationEvent

	TaskID string `json:"taskId"`
	UserID string `json:"userId"`
}

var _ events.IntegrationEvent = (*TaskAssignedEvent)(nil)

const TaskAssignedEventName = "task.assigned"

func NewTaskAssignedEvent(correlationID string, taskID string, userID string) (*TaskAssignedEvent, error) {
	event, err := v1.NewEvent(TaskAssignedEventName)
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

	return &TaskAssignedEvent{
		IntegrationEvent: ie,
		TaskID:           taskID,
		UserID:           userID,
	}, nil
}

func (e TaskAssignedEvent) ID() string {
	return e.IntegrationEvent.ID()
}

func (e TaskAssignedEvent) Name() string {
	return e.IntegrationEvent.Name()
}

func (e TaskAssignedEvent) CorrelationID() string {
	return e.IntegrationEvent.CorrelationID()
}

func (e TaskAssignedEvent) OccurredAt() time.Time {
	return e.IntegrationEvent.OccurredAt()
}

func (e TaskAssignedEvent) MarshalJSON() ([]byte, error) {
	payload := struct {
		IntegrationEvent *v1.IntegrationEvent `json:"integrationEvent"`
		UserID           string               `json:"userId"`
		TaskID           string               `json:"taskId"`
	}{
		IntegrationEvent: e.IntegrationEvent,
		UserID:           e.UserID,
		TaskID:           e.TaskID,
	}

	return json.Marshal(payload)
}

func (e *TaskAssignedEvent) UnmarshalJSON(data []byte) error {
	payload := struct {
		IntegrationEvent *v1.IntegrationEvent `json:"integrationEvent"`
		UserID           string               `json:"userId"`
		TaskID           string               `json:"taskId"`
	}{}

	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}

	e.IntegrationEvent = payload.IntegrationEvent
	e.UserID = payload.UserID
	e.TaskID = payload.TaskID
	return nil
}
