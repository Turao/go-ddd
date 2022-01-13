package task

import "errors"

type TaskID = string

type Task struct {
	ID TaskID `json:"id"`
}

var (
	ErrInvalidTaskID = errors.New("invalid task id")
)

func NewTask(id TaskID) (*Task, error) {
	if id == "" {
		return nil, ErrInvalidTaskID
	}

	return &Task{
		ID: id,
	}, nil
}
