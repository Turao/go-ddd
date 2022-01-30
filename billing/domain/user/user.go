package user

import (
	"errors"

	"github.com/turao/go-ddd/tasks/domain/task"
	"github.com/turao/go-ddd/users/domain/user"
)

type UserID = user.UserID
type TaskID = task.TaskID

type User struct {
	ID    UserID          `json:"userId"`
	Tasks map[TaskID]bool `json:"tasks"`
}

var (
	ErrInvalidUserID = errors.New("invalid user id")
	ErrInvalidTaskID = errors.New("invalid task id")
)

func NewUser(userID string) (*User, error) {
	if userID == "" {
		return nil, ErrInvalidUserID
	}

	return &User{
		ID:    userID,
		Tasks: make(map[TaskID]bool),
	}, nil
}

func (u *User) AddTask(taskID TaskID) error {
	if taskID == "" {
		return ErrInvalidTaskID
	}

	u.Tasks[taskID] = true
	return nil
}

func (u *User) RemoveTask(taskID TaskID) error {
	if taskID == "" {
		return ErrInvalidTaskID
	}

	delete(u.Tasks, taskID)
	return nil
}
