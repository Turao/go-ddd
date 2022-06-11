package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/turao/go-ddd/projects/application"
)

func (a Application) HandleListProjects(rw http.ResponseWriter, r *http.Request) {
	req := application.ListProjectsQuery{}
	res, err := a.Delegate.Queries.ListProjects.Handle(context.Background(), req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(res)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(payload)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
