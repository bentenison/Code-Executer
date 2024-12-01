package kafkaconsumer

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Consumer wraps Kafka consumer configuration and methods
type Consumer struct {
	consumer *kafka.Consumer
}

// NewConsumer creates a new Kafka consumer instance
func NewConsumer(broker, groupID, topic string) (*Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,     // Kafka broker
		"group.id":          groupID,    // Consumer group ID
		"auto.offset.reset": "earliest", // Read from the earliest offset
	})
	if err != nil {
		return nil, err
	}

	// Subscribe to the topic
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{consumer: consumer}, nil
}

// ConsumeMessages continuously consumes messages from the Kafka topic
func (c *Consumer) ConsumeMessages() ([]string, error) {
	var messages []string
	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error consuming message: %s", err)
			return nil, err
		}

		messages = append(messages, string(msg.Value))
	}

	return messages, nil
}

// Close gracefully shuts down the consumer
func (c *Consumer) Close() {
	if c.consumer != nil {
		c.consumer.Close()
	}
}
