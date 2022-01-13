package user

import "errors"

type UserID = string

type User struct {
	ID UserID `json:"id"`
}

var (
	ErrInvalidUserID = errors.New("invalid user id")
)

func NewUser(id UserID) (*User, error) {
	if id == "" {
		return nil, ErrInvalidUserID
	}

	return &User{
		ID: id,
	}, nil
}
