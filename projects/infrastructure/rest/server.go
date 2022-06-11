package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/turao/go-ddd/projects/application"
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
	router.HandleFunc("/project", app.HandleCreateProject).Methods(http.MethodPost)
	router.HandleFunc("/project", app.HandleDeleteProject).Methods(http.MethodDelete)
	router.HandleFunc("/project", app.HandleUpdateProject).Methods(http.MethodPut)
	router.HandleFunc("/project/{projectId}", app.HandleFindProject).Methods(http.MethodGet)
	router.HandleFunc("/projects", app.HandleListProjects).Methods(http.MethodGet)

	// prepare HTTP server
	httpServer := &http.Server{
		Addr:         "0.0.0.0:8085",
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
