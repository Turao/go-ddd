package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/turao/go-ddd/projects/application"
)

func (a Application) HandleUpdateProject(rw http.ResponseWriter, r *http.Request) {
	var req application.UpdateProjectCommand
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = a.Delegate.Commands.UpdateProject.Handle(context.Background(), req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
