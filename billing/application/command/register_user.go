package command

import (
	"context"

	"github.com/turao/go-ddd/billing/application"
)

type RegisterUserCommandHandler struct {
}

var _ application.RegisterUserCommandHandler = (*RegisterUserCommandHandler)(nil)

func NewRegisterUserCommandHandler() (*RegisterUserCommandHandler, error) {
	return &RegisterUserCommandHandler{}, nil
}

func (h RegisterUserCommandHandler) Handle(ctx context.Context, req application.RegisterUserCommand) error {
	return nil
}
