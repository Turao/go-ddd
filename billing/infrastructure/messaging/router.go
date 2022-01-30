package messaging

import (
	"context"
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
}

func (r *Router) Init() error {
	logger := watermill.NewStdLogger(true, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return err
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)

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

	r.router = router
	return nil
}

func (r Router) Run(ctx context.Context) error {
	return r.router.Run(ctx)
}

func (r Router) Close() error {
	return r.router.Close()
}
