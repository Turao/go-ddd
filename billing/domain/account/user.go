package account

import (
	"errors"

	"github.com/turao/go-ddd/users/domain/user"
)

type UserID = user.UserID

type User struct {
	ID UserID `json:"id"`
}

var (
	ErrInvalidUserID = errors.New("invalid user id")
)

func NewUser(userID UserID) (*User, error) {
	if userID == "" {
		return nil, ErrInvalidUserID
	}

	return &User{
		ID: userID,
	}, nil
}
