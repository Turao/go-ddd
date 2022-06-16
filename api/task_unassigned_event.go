package api

import (
	"context"
	"encoding/json"
	"errors"
	"time"

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

func (e TaskUnassignedEvent) ID() string {
	return e.IntegrationEvent.ID()
}

func (e TaskUnassignedEvent) Name() string {
	return e.IntegrationEvent.Name()
}

func (e TaskUnassignedEvent) CorrelationID() string {
	return e.IntegrationEvent.CorrelationID()
}

func (e TaskUnassignedEvent) OccurredAt() time.Time {
	return e.IntegrationEvent.OccurredAt()
}

func (e TaskUnassignedEvent) MarshalJSON() ([]byte, error) {
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

func (e *TaskUnassignedEvent) UnmarshalJSON(data []byte) error {
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
