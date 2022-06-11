package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/turao/go-ddd/tasks/application"
)

func (a Application) HandleCreateTask(rw http.ResponseWriter, r *http.Request) {
	var req application.CreateTaskCommand
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = a.Delegate.Commands.CreateTaskCommand.Handle(context.Background(), req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
