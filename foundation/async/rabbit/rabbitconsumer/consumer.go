package rabbitconsumer

import (
	"github.com/streadway/amqp"
)

type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewConsumer(rabbitmqURL, queueName string) (*Consumer, error) {
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

	// Ensure the queue exists (durable, non-exclusive)
	_, err = ch.QueueDeclare(
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

	return &Consumer{
		connection: conn,
		channel:    ch,
	}, nil
}

func (c *Consumer) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	// Consume messages from the queue
	msgs, err := c.channel.Consume(
		queueName, // queue name
		"",        // consumer tag
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

// Close gracefully shuts down the consumer connection and channel
func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.connection != nil {
		c.connection.Close()
	}
}

// Define RabbitMQ connection URL
// rabbitmqURL := "amqp://guest:guest@localhost:5672/" // Default URL

// // Initialize Producer
// producer, err := rabbitmqproducer.NewProducer(rabbitmqURL)
// if err != nil {
// 	log.Fatalf("Error creating producer: %s", err)
// }
// defer producer.Close()

// // Produce a message to a RabbitMQ queue
// queueName := "test-queue"
// message := "Hello RabbitMQ from Go!"
// err = producer.ProduceMessage(queueName, message)
// if err != nil {
// 	log.Fatalf("Error producing message: %s", err)
// }
// fmt.Println("Message sent successfully!")

// // Initialize Consumer
// consumer, err := rabbitmqconsumer.NewConsumer(rabbitmqURL, queueName)
// if err != nil {
// 	log.Fatalf("Error creating consumer: %s", err)
// }
// defer consumer.Close()

// // Consume messages
// fmt.Println("Consuming messages...")
// msgs, err := consumer.ConsumeMessages(queueName)
// if err != nil {
// 	log.Fatalf("Error consuming messages: %s", err)
// }

// // Listen for messages
// for msg := range msgs {
// 	fmt.Printf("Received message: %s\n", msg.Body)
// }

// // Sleep for a while before closing the consumer
// time.Sleep(2 * time.Second)

// producer, err := kafkaproducer.NewProducer("localhost:9092")
// 	if err != nil {
// 		log.Fatalf("Error creating producer: %s", err)
// 	}
// 	defer producer.Close()

// 	err = producer.ProduceMessage("my-topic", "Hello Kafka!")
// 	if err != nil {
// 		log.Fatalf("Error producing message: %s", err)
// 	}

// 	// Consumer Example
// 	consumer, err := kafkaconsumer.NewConsumer("localhost:9092", "my-group", "my-topic")
// 	if err != nil {
// 		log.Fatalf("Error creating consumer: %s", err)
// 	}
// 	defer consumer.Close()

// 	messages, err := consumer.ConsumeMessages()
// 	if err != nil {
// 		log.Fatalf("Error consuming messages: %s", err)
// 	}

// 	for _, msg := range messages {
// 		fmt.Println("Consumed message:", msg)
// 	}
