package rest

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/turao/go-ddd/users/application"
)

type mockRegisterUserCommandHandler struct{}

func (h mockRegisterUserCommandHandler) Handle(ctx context.Context, cmd application.RegisterUserCommand) error {
	if cmd.Username == "" {
		return errors.New("'mock' validation error")
	}
	return nil
}

type mockListUsersQueryHandler struct{}

func (h mockListUsersQueryHandler) Handle(ctx context.Context, cmd application.ListUsersQuery) (*application.ListUsersResponse, error) {
	return nil, nil
}

func newTestApplication() *Application {

	app := &Application{
		Delegate: &application.Application{
			Commands: application.Commands{
				RegisterUserCommand: mockRegisterUserCommandHandler{},
			},
			Queries: application.Queries{
				ListUsersQuery: mockListUsersQueryHandler{},
			},
		},
	}
	return app
}

func TestRegisterUserHandler(t *testing.T) {
	type test struct {
		Body           string
		ExpectedStatus int
	}

	app := newTestApplication()

	tests := map[string]test{
		"valid payload":      {Body: `{ "username": "dummy" }`, ExpectedStatus: http.StatusCreated},
		"unprocessable json": {Body: `{ "usern }`, ExpectedStatus: http.StatusBadRequest},
		"empty body":         {Body: `{}`, ExpectedStatus: http.StatusInternalServerError},
	}

	for name, test := range tests {
		req := httptest.NewRequest("POST", "/users", strings.NewReader(test.Body))
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(app.HandleRegisterUser)

		handler.ServeHTTP(recorder, req)

		assert.Equalf(t, recorder.Code, test.ExpectedStatus, name)
	}

}
