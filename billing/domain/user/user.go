package user

import (
	"errors"

	"github.com/turao/go-ddd/tasks/domain/task"
	"github.com/turao/go-ddd/users/domain/user"
)

type UserID = user.UserID
type TaskID = task.TaskID

type User struct {
	ID    UserID   `json:"userId"`
	Tasks []TaskID `json:"tasks"`
}

var (
	ErrInvalidUserID = errors.New("invalid user id")
)

func NewUser(userID string) (*User, error) {
	if userID == "" {
		return nil, ErrInvalidUserID
	}

	return &User{
		ID:    userID,
		Tasks: make([]string, 0),
	}, nil
}
