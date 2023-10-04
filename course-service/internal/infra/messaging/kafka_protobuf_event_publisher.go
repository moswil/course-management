package messaging

import (
	"context"
	"fmt"
	"log"

	"github.com/moswil/course-management/internal/core/domain/event/pb"
	"github.com/moswil/course-management/internal/core/port/event"
	"github.com/moswil/course-management/internal/infra/serialization"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type KafkaConfig struct {
	Topic   string   `mapstructure:"topic"`
	Brokers []string `mapstructure:"brokers"`
}

// KafkaEventPublisher publishes events to a Kafka topic.
type KafkaProtobufEventPublisher struct {
	writer     *kafka.Writer
	serializer *serialization.ProtobufEventSerializer // Inject the serializer
}

// NewKafkaEventPublisher creates a new KafkaEventPublisher instance.
func NewKafkaProtobufEventPublisher() event.EventPublisher {
	kafkaConfig := KafkaConfig{
		Topic:   viper.GetString("messaging.kafka.topic"),
		Brokers: viper.GetStringSlice("messaging.kafka.brokers"),
	}
	// if err := viper.Unmarshal(&kafkaConfig); err != nil {
	// 	log.Printf("error loading Kafka Config: %v\n", err.Error())
	// }

	log.Printf("Kafka Config: %v\n", kafkaConfig)
	log.Printf("Kafka Config: %v\n", viper.AllKeys())

	// writer := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{"localhost:9092"}, // Kafka broker address
	// 	Topic:   "course-events",            // Replace with your actual Kafka topic name
	// })

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: kafkaConfig.Brokers, // Kafka broker address
		Topic:   kafkaConfig.Topic,   // Replace with your actual Kafka topic name
	})

	return &KafkaProtobufEventPublisher{
		writer:     writer,
		serializer: serialization.NewProtobufEventSerializer(), // Replace with your actual event serializer
	}
}

// Publish serializes and publishes an event to the Kafka topic.
func (p *KafkaProtobufEventPublisher) Publish(ctx context.Context, evt interface{}) error {
	// Serialize the event to Protobuf format
	log.Printf("event received: %v", evt)

	// Assuming event is of type interface{}
	event, ok := evt.(*pb.CourseCreatedEvent)
	if !ok {
		// Handle the case where event is not a *pb.CourseCreatedEvent
		return fmt.Errorf("event is not a CourseCreatedEvent")
	}

	// event before serialization
	log.Printf("event before serialization: %v - %v - %v", event.CourseId, event.Title, event.Instructor)

	// Now you can call the SerializeCourseCreatedEvent function
	serializedEvent, err := p.serializer.SerializeCourseCreatedEvent(event)
	if err != nil {
		// Handle the error
		return err
	}

	// Publish the serialized event to the Kafka topic
	message := kafka.Message{
		Value: serializedEvent,
	}

	err = p.writer.WriteMessages(ctx, message)
	if err != nil {
		log.Printf("Error publishing message: %v", err)
		return err
	}

	return nil
}
