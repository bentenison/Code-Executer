package rabbitproducer

import (
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewProducer(rabbitmqURL string) (*Producer, error) {
	// Establish connection to RabbitMQ server
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	// Create a new channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Producer{
		connection: conn,
		channel:    ch,
	}, nil
}

func (p *Producer) ProduceMessage(queueName string, message string) error {
	// Ensure the queue exists (durable, non-exclusive)
	_, err := p.channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
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
			ContentType: "text/json",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Message sent to queue %s: %s", queueName, message)
	return nil
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
