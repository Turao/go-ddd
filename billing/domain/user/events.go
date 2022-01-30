package user

import "github.com/turao/go-ddd/events"

type UserRegisteredEvent struct {
	events.DomainEvent `json:"domainEvent"`
}

// var (
// 	ErrInvalidUserID   = errors.New("invalid user id")
// 	ErrInvalidTaskID   = errors.New("invalid task id")
// )

func NewUserRegisteredEvent(userID string) (*UserRegisteredEvent, error) {
	domainEvent, err := events.NewDomainEvent("user.registered", userID)
	if err != nil {
		return nil, err
	}

	return &UserRegisteredEvent{
		DomainEvent: domainEvent,
	}, nil
}

type TaskAssignedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	TaskID             TaskID `json:"taskId"`
}

func NewTaskAssignedEvent(userID string, taskID string) (*TaskAssignedEvent, error) {
	domainEvent, err := events.NewDomainEvent("task.assigned", userID)
	if err != nil {
		return nil, err
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	return &TaskAssignedEvent{
		DomainEvent: domainEvent,
		TaskID:      taskID,
	}, nil
}

type TaskUnassignedEvent struct {
	events.DomainEvent `json:"domainEvent"`
	TaskID             TaskID `json:"taskId"`
}

func NewTaskUnassignedEvent(userID string, taskID string) (*TaskUnassignedEvent, error) {
	domainEvent, err := events.NewDomainEvent("task.unassigned", userID)
	if err != nil {
		return nil, err
	}

	if taskID == "" {
		return nil, ErrInvalidTaskID
	}

	return &TaskUnassignedEvent{
		DomainEvent: domainEvent,
		TaskID:      taskID,
	}, nil
}
