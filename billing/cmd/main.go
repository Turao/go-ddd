package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	watermillAMQP "github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"

	"github.com/turao/go-ddd/api/amqp"
)

func main() {
	queue := watermillAMQP.NewDurableQueueConfig("amqp://localhost:5672")
	logger := watermill.NewStdLogger(false, false)

	subscriber, err := watermillAMQP.NewSubscriber(queue, logger)
	if err != nil {
		log.Fatalln(err)
	}
	defer subscriber.Close()

	ures, err := amqp.NewUserRegisteredEventSubscriber(subscriber)
	if err != nil {
		log.Fatalln(err)
	}

	events, err := ures.Subscribe(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for event := range events {
			d, err := json.MarshalIndent(event, "", " ")
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("received event:", string(d))
		}
	}()

	time.Sleep(30 * time.Second)
}
