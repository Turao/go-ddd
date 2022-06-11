package rest

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turao/go-ddd/tasks/application"
)

func (a Application) HandleUnassignUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req := application.UnassignUserCommand{
		TaskID: vars["taskId"],
	}

	err := a.Delegate.Commands.UnassignUserCommand.Handle(context.Background(), req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
