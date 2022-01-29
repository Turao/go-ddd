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
	queueConfig := watermillAMQP.NewDurableQueueConfig("amqp://localhost:5672")
	logger := watermill.NewStdLogger(false, false)

	subscriber, err := watermillAMQP.NewSubscriber(queueConfig, logger)
	if err != nil {
		log.Fatalln(err)
	}
	defer subscriber.Close()

	ures, err := amqp.NewUserRegisteredEventSubscriber(subscriber)
	if err != nil {
		log.Fatalln(err)
	}

	userRegisteredEvents, err := ures.Subscribe(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for event := range userRegisteredEvents {
			d, err := json.MarshalIndent(event, "", " ")
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("received event:", string(d))
		}
	}()

	tsues, err := amqp.NewTaskStatusUpdatedEventSubscriber(subscriber)
	if err != nil {
		log.Fatalln(err)
	}

	taskStatusUpdatedEvents, err := tsues.Subscribe(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for event := range taskStatusUpdatedEvents {
			d, err := json.MarshalIndent(event, "", " ")
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("received event:", string(d))
		}
	}()

	taes, err := amqp.NewTaskAssignedEventSubscriber(subscriber)
	if err != nil {
		log.Fatalln(err)
	}

	taskAssignedEvents, err := taes.Subscribe(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for event := range taskAssignedEvents {
			d, err := json.MarshalIndent(event, "", " ")
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("received event:", string(d))
		}
	}()

	tues, err := amqp.NewTaskUnassignedEventSubscriber(subscriber)
	if err != nil {
		log.Fatalln(err)
	}

	taskUnassignedEvents, err := tues.Subscribe(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for event := range taskUnassignedEvents {
			d, err := json.MarshalIndent(event, "", " ")
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("received event:", string(d))
		}
	}()

	time.Sleep(60 * time.Second)
}
