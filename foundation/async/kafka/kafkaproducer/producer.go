// kafkaproducer/producer.go

package kafkaproducer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Producer wraps Kafka producer configuration and methods
type Producer struct {
	producer *kafka.Producer
}

// NewProducer creates a new Kafka producer instance
func NewProducer(broker string) (*Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker, // Kafka broker
	})
	if err != nil {
		return nil, err
	}

	return &Producer{producer: producer}, nil
}

// ProduceMessage sends a message to the specified Kafka topic
func (p *Producer) ProduceMessage(topic string, message string) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(message),
	}

	// Produce message asynchronously
	p.producer.Produce(msg, nil)

	// Wait for all messages to be delivered before returning
	p.producer.Flush(15 * 1000)

	return nil
}

// Close gracefully shuts down the producer
func (p *Producer) Close() {
	if p.producer != nil {
		p.producer.Close()
	}
}
