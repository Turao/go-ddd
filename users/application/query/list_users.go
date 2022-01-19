package query

import (
	"context"

	"github.com/turao/go-ddd/users/application"
	"github.com/turao/go-ddd/users/domain/user"
)

type ListUsersQueryHandler struct {
	repository user.Repository
}

func NewListUsersQueryHandler(repository user.Repository) *ListUsersQueryHandler {
	return &ListUsersQueryHandler{
		repository: repository,
	}
}

func (h ListUsersQueryHandler) Handle(ctx context.Context, req application.ListUsersQuery) (*application.ListUsersResponse, error) {
	us, err := h.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	usDTO := make([]application.User, 0)
	for _, u := range us {
		usDTO = append(usDTO, application.User{
			UserID:   u.ID,
			Username: u.Name,
		})
	}

	return &application.ListUsersResponse{
		Users: usDTO,
	}, nil
}
