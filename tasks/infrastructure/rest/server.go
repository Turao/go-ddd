package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/turao/go-ddd/tasks/application"
)

type Server struct {
	HTTPServer *http.Server
	App        *Application
}

// Application is a thin layer that decorates another Application,
// unmarshalling HTTP requests, handling commands/queries, and marshalling responses to HTTP format
type Application struct {
	Delegate *application.Application
}

func NewServer(app *Application) (*Server, error) {
	// prepare application routes
	router := mux.NewRouter()
	router.Use(ContentTypeJSON)
	router.HandleFunc("/task", app.HandleCreateTask).Methods(http.MethodPost)
	router.HandleFunc("/task/{taskId}/assign", app.HandleAssignToUser).Methods(http.MethodPost)
	router.HandleFunc("/task/{taskId}/unassign", app.HandleUnassignUser).Methods(http.MethodPost)
	router.HandleFunc("/tasks/project/{projectId}", app.HandleTasksByProject).Methods(http.MethodGet)
	router.HandleFunc("/tasks/user/{userId}", app.HandleTasksByAssignedUser).Methods(http.MethodGet)

	// prepare HTTP server
	httpServer := &http.Server{
		Addr:         "0.0.0.0:8086",
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
