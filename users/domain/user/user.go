package user

import "errors"

type UserID = string

type User struct {
	ID UserID `json:"id"`

	Name string `json:"name"`
}

var (
	ErrEmptyUserID   = errors.New("empty user id")
	ErrEmptyUserName = errors.New("invalid user name")
)

func NewUser(id UserID, name string) (*User, error) {
	if id == "" {
		return nil, ErrEmptyUserID
	}

	if name == "" {
		return nil, ErrEmptyUserName
	}

	return &User{
		ID:   id,
		Name: name,
	}, nil
}
