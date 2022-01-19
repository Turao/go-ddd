package user

import "errors"

type UserID = string

type User struct {
	ID UserID `json:"id"`

	Name string `json:"name"`
}

var (
	ErrInvalidUserID   = errors.New("invalid user id")
	ErrInvalidUserName = errors.New("invalid user name")
)

func NewUser(id UserID, name string) (*User, error) {
	if id == "" {
		return nil, ErrInvalidUserID
	}

	if name == "" {
		return nil, ErrInvalidUserName
	}

	return &User{
		ID:   id,
		Name: name,
	}, nil
}
