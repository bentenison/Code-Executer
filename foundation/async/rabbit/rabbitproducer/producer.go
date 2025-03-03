package rabbitproducer

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/streadway/amqp"
)

const (
	maxRetries        = 5
	initialBackoff    = 100 * time.Millisecond
	backoffMultiplier = 2
)

// Producer struct holds the connection and channel
type Producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queues     []string
}

// NewProducer establishes a connection and declares multiple queues
func NewProducer(rabbitmqURL string, queueNames []string) (*Producer, error) {
	var conn *amqp.Connection
	var err error

	// Exponential backoff for connection retries
	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(rabbitmqURL)
		if err == nil {
			break
		}

		// Calculate backoff time
		backoffTime := time.Duration(math.Pow(float64(backoffMultiplier), float64(i))) * initialBackoff
		// Log the error and retry
		fmt.Printf("Failed to connect to RabbitMQ: %v. Retrying in %v...\n", err, backoffTime)
		time.Sleep(backoffTime)
	}

	if err != nil {
		return nil, err
	}

	// Create a new channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Declare each queue
	for _, queueName := range queueNames {
		_, err := ch.QueueDeclare(
			queueName, // queue name
			true,      // durable
			false,     // auto-delete
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if err != nil {
			return nil, err
		}
	}

	return &Producer{
		connection: conn,
		channel:    ch,
		queues:     queueNames,
	}, nil
}

// ProduceMessage sends a message to a specified queue
func (p *Producer) ProduceMessage(queueName string, message interface{}) error {
	// Check if the queue exists in the producer's list
	for _, q := range p.queues {
		if q == queueName {
			jsonData, err := json.Marshal(message)
			if err != nil {
				return err
			}
			// Publish the message to the queue
			err = p.channel.Publish(
				"",        // exchange
				queueName, // routing key (queue name)
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType: "application/json", // Changed to plain text
					Body:        jsonData,
				},
			)
			if err != nil {
				return err
			}

			log.Printf("Message sent to queue %s: %s", queueName, message)
			return nil
		}
	}

	return log.Output(1, "Queue does not exist: "+queueName)
}

// Close gracefully shuts down the producer connection and channel
func (p *Producer) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.connection != nil {
		p.connection.Close()
	}
}
