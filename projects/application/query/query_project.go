package query

import (
	"context"
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
	repo       project.Repository
	eventStore events.EventStore
}

func NewFindProjectQueryHandler(repo project.Repository, es events.EventStore) *FindProjectHandler {
	return &FindProjectHandler{
		repo:       repo,
		eventStore: es,
	}
}

func (h *FindProjectHandler) Handle(ctx context.Context, req FindProjectQuery) (*FindProjectResponse, error) {
	var p project.ProjectAggregate
	evts := h.eventStore.FilterByAggregateID(req.ID)
	for _, evt := range evts {
		log.Println(reflect.TypeOf(evt))

		devt := evt.(events.DomainEvent)
		if devt.AggregateID() == req.ID {
			if err := p.HandleEvent(devt); err != nil {
				return nil, err
			}
		}
	}
	return &FindProjectResponse{
		ID:     p.Project.ID,
		Name:   p.Project.Name,
		Active: p.Project.Active,
	}, nil
}
