package messaging

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/turao/go-ddd/api"
	"github.com/turao/go-ddd/billing/application"
)

type Router struct {
	router *message.Router

	CreateAccountCommandHandler      application.CreateAccountCommandHandler
	AddTaskToUserCommandHandler      application.AddTaskToUserCommandHandler
	RemoveTaskFromUserCommandHandler application.RemoveTaskFromUserCommandHandler
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

func ErrorLogger(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) ([]*message.Message, error) {
		msgs, err := h(msg)
		if err != nil {
			log.Println("[ERROR]:", err.Error())
		}
		return msgs, err
	}
}

func (r *Router) Init() error {
	logger := watermill.NewStdLogger(true, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return err
	}

	// queueConfig := amqp.NewDurableQueueConfig("amqp://localhost:5672")
	subscriberConfig := kafka.SubscriberConfig{
		Brokers:       []string{"localhost:29092"},
		Unmarshaler:   kafka.DefaultMarshaler{},
		ConsumerGroup: "billing",
	}
	subscriber, err := kafka.NewSubscriber(subscriberConfig, logger)
	if err != nil {
		return err
	}

	publisherConfig := kafka.PublisherConfig{
		Brokers:   []string{"localhost:29092"},
		Marshaler: kafka.DefaultMarshaler{},
	}
	publisher, err := kafka.NewPublisher(publisherConfig, logger)
	if err != nil {
		return err
	}

	poisonMiddleware, err := middleware.PoisonQueue(publisher, "unprocessable.messages")
	if err != nil {
		return err
	}

	router.AddNoPublisherHandler(
		"billing.user.registered.handler",
		api.UserRegisteredEventName,
		subscriber,
		UserRegisteredEventHandler{
			CommandHandler: r.CreateAccountCommandHandler,
		}.Handle,
	)

	router.AddNoPublisherHandler(
		"billing.task.assigned.handler",
		api.TaskAssignedEventName,
		subscriber,
		TaskAssignedEventHandler{
			CommandHandler: r.AddTaskToUserCommandHandler,
		}.Handle,
	)

	router.AddNoPublisherHandler(
		"billing.task.unassigned.handler",
		api.TaskUnassignedEventName,
		subscriber,
		TaskUnassignedEventHandler{
			CommandHandler: r.RemoveTaskFromUserCommandHandler,
		}.Handle,
	)

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		middleware.Retry{
			MaxRetries: 1,
		}.Middleware,
		// middleware.Recoverer,
		poisonMiddleware,
		MessageLogger,
		// ErrorLogger,
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
