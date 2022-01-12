package query

import (
	"context"
	"errors"
	"log"
	"reflect"

	"github.com/turao/go-ddd/events"
	"github.com/turao/go-ddd/projects/domain/project"
)

type FindProjectQuery struct {
	ID string `json:"id"`
}

type FindProjectResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type FindProjectHandler struct {
	eventStore events.EventStore
}

func NewFindProjectQueryHandler(es events.EventStore) *FindProjectHandler {
	return &FindProjectHandler{
		eventStore: es,
	}
}

func (h *FindProjectHandler) Handle(ctx context.Context, req FindProjectQuery) (*FindProjectResponse, error) {
	var p project.ProjectAggregate
	evts, err := h.eventStore.Events(context.Background())
	if err != nil {
		return nil, err
	}

	log.Println("events", evts)

	for _, evt := range evts {
		log.Println(reflect.TypeOf(evt))

		devt := evt.(events.DomainEvent)
		if devt.AggregateID() == req.ID {
			if err := p.HandleEvent(devt); err != nil {
				return nil, err
			}
		}
	}

	if p.Project == nil {
		return nil, errors.New("cannot reconstruct project from events")
	}

	return &FindProjectResponse{
		ID:     p.Project.ID,
		Name:   p.Project.Name,
		Active: p.Project.Active,
	}, nil
}
