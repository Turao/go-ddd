package account

import "github.com/turao/go-ddd/users/domain/user"

type CreateAccountCommand struct {
	UserID user.UserID
}

type AddTaskToUserCommand struct {
	TaskID TaskID
}

type RemoveTaskFromUserCommand struct {
	TaskID TaskID
}
