package account

import "github.com/turao/go-ddd/users/domain/user"

type CreateAccountCommand struct {
	UserID user.UserID
}

type AddTaskCommand struct {
	TaskID TaskID
}

type RemoveTaskCommand struct {
	TaskID TaskID
}
