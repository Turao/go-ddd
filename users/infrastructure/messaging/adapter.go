package messaging

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type Adapter struct {
	OnRegisterUser OnRegisterUser
}

func NewAdapter(ru OnRegisterUser) (*Adapter, error) {
	return &Adapter{
		OnRegisterUser: ru,
	}, nil
}

func (a Adapter) RegisterHandlers() error {
	logger := watermill.NewStdLogger(false, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return err
	}

	channel := gochannel.NewGoChannel(gochannel.Config{}, logger)

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Recoverer,
	)

	router.AddHandler(
		"onRegisterUser",
		"register.user.request",
		channel,
		"user.registered",
		channel,
		a.OnRegisterUser.Handle,
	)

	if err := router.Run(context.Background()); err != nil {
		return err
	}

	return nil
}
