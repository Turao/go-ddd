package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turao/go-ddd/tasks/application"
)

func (a Application) HandleAssignToUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var body struct {
		UserID string `json:"userId"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	defer r.Body.Close()

	req := application.AssignToUserCommand{
		TaskID: vars["taskId"],
		UserID: body.UserID,
	}

	err = a.Delegate.Commands.AssignToUserCommand.Handle(context.Background(), req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
