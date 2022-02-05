package infrastructure

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
		Delegate: application.Application{
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

	tests := []test{
		{Body: `{ "username": "dummy" }`, ExpectedStatus: http.StatusCreated},
		{Body: `{ "usern }`, ExpectedStatus: http.StatusBadRequest},
		{Body: `{}`, ExpectedStatus: http.StatusInternalServerError},
	}

	for _, test := range tests {
		req := httptest.NewRequest("POST", "/users", strings.NewReader(test.Body))

		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(app.HandleRegisterUser)

		handler.ServeHTTP(recorder, req)

		if status := recorder.Code; status != test.ExpectedStatus {
			t.Errorf("handler returned wrong http code: got %v - wanted %v", status, test.ExpectedStatus)
		}
	}

}
