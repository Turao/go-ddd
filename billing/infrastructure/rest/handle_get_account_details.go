package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turao/go-ddd/billing/application"
)

func (a Application) HandleGetAccountDetails(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req := application.GetAccountDetailsQuery{
		AccountID: vars["accountId"],
	}

	res, err := a.Delegate.Queries.GetAccountDetails.Handle(context.Background(), req)
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
