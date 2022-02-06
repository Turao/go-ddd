package infrastructure

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/turao/go-ddd/tasks/application"
)

type Server struct {
	HTTPServer *http.Server
	App        *Application
}

func NewServer(app *Application) (*Server, error) {
	// prepare application routes
	router := mux.NewRouter()
	router.Use(ContentTypeJSON)
	router.HandleFunc("/task", app.HandleCreateTask).Methods(http.MethodPost)
	router.HandleFunc("/project/{projectId}", app.HandleTasksByProject).Methods(http.MethodGet)
	router.HandleFunc("/user/{userId}", app.HandleTasksByAssignedUser).Methods(http.MethodGet)

	// prepare HTTP server
	httpServer := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      router,
	}

	return &Server{
		HTTPServer: httpServer,
		App:        app,
	}, nil
}

func (s Server) ListenAndServe() error {
	return s.HTTPServer.ListenAndServe()
}

func (s Server) Close() error {
	return s.HTTPServer.Close()
}

func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Application is a thin layer that decorates another Application,
// unmarshalling HTTP requests, handling commands/queries, and marshalling responses to HTTP format
type Application struct {
	Delegate *application.Application
}

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

func (a Application) HandleTasksByAssignedUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req := application.TasksByAssignedUserQuery{
		UserID: vars["userId"],
	}

	res, err := a.Delegate.Queries.TasksByAssignedUserQuery.Handle(context.Background(), req)
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

func (a Application) HandleTasksByProject(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req := application.TasksByProjectQuery{
		ProjectID: vars["projectId"],
	}

	res, err := a.Delegate.Queries.TasksByProjectQuery.Handle(context.Background(), req)
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