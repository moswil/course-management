package messaging

import (
	"context"
	"encoding/json"

	"github.com/moswil/course-management/internal/core/interface/event"
	"github.com/segmentio/kafka-go"
)

type KafkaEventPublisher struct {
	// Initialize Kafka configuration and connections here
	writer *kafka.Writer
}

func NewKafkaEventPublisher() event.EventPublisher {
	// Initialize and configure Kafka connections
	return &KafkaEventPublisher{
		// Initialize Kafka writer
		writer: &kafka.Writer{
			Addr:  kafka.TCP("localhost:9092"), // Kafka broker address
			Topic: "course-events",             // Kafka topic for course events
		},
	}
}

func (p *KafkaEventPublisher) Publish(ctx context.Context, event interface{}) error {
	// log.Printf("course-event-received: %v\n", event)
	// Serialize the event to JSON or your desired format
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Publish the event to the Kafka topic
	if err := p.writer.WriteMessages(
		ctx,
		// Pass the serialized event data here
		kafka.Message{Value: eventJSON},
	); err != nil {
		return err
	}
	return nil
}
