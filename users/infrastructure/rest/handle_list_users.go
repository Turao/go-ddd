package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/turao/go-ddd/users/application"
)

func (a Application) HandleListUsers(rw http.ResponseWriter, r *http.Request) {
	var req application.ListUsersQuery
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res, err := a.Delegate.Queries.ListUsersQuery.Handle(context.Background(), req)
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
