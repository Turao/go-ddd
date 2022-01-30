package messaging

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type Router struct {
	router *message.Router

	RegisterUserCommandHandler application.RegisterUserCommandHandler
	AssignTaskCommandHandler   application.AssignTaskCommandHandler
	UnassignTaskCommandHandler application.UnassignTaskCommandHandler
}

func MessageLogger(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) ([]*message.Message, error) {
		d, err := json.Marshal(msg.Payload)
		if err != nil {
			return nil, err
		}
		log.Println("received message: ", string(d))
		defer log.Println("message processed!")
		return h(msg)
	}
}

func (r *Router) Init() error {
	logger := watermill.NewStdLogger(true, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return err
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		middleware.Retry{
			MaxRetries: 1,
		}.Middleware,
		middleware.Recoverer,
		MessageLogger,
	)

	queueConfig := amqp.NewDurableQueueConfig("amqp://localhost:5672")
	subscriber, err := amqp.NewSubscriber(queueConfig, logger)
	if err != nil {
		log.Fatalln(err)
	}

	router.AddNoPublisherHandler(
		"register.user",
		api.UserRegisteredEventName,
		subscriber,
		RegisterUserCommandHandler{
			CommandHandler: r.RegisterUserCommandHandler,
		}.Handle,
	)

	router.AddNoPublisherHandler(
		"assign.task",
		api.TaskAssignedEventName,
		subscriber,
		AssignTaskCommandHandler{
			CommandHandler: r.AssignTaskCommandHandler,
		}.Handle,
	)

	router.AddNoPublisherHandler(
		"unassign.task",
		api.TaskAssignedEventName,
		subscriber,
		UnassignTaskCommandHandler{
			CommandHandler: r.UnassignTaskCommandHandler,
		}.Handle,
	)

	r.router = router
	return nil
}

func (r Router) Run(ctx context.Context) error {
	return r.router.Run(ctx)
}

func (r Router) Close() error {
	return r.router.Close()
}
