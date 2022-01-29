package domain

import "errors"

type UserID = string

type User struct {
	ID UserID
}

var (
	ErrInvalidUserID = errors.New("invalid user id")
)

func NewUser(userID string) (*User, error) {
	if userID == "" {
		return nil, ErrInvalidUserID
	}

	return &User{
		ID: userID,
	}, nil
}
