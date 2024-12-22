package rabbitconsumer

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bentenison/microservice/foundation/logger"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/streadway/amqp"
)

type DataDAO interface {
	StoreCodeExecutionStatsES(ctx context.Context, codeStats []byte) error
	StoreChallengeDataES(ctx context.Context, challengeData []byte) error
	StorePerformanceDataES(ctx context.Context, performanceData []byte) error
}

// Consumer struct holds the connection and channel
type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queues     []string
	es         *elasticsearch.Client
	log        *logger.CustomLogger
}

// NewConsumer establishes a connection and declares multiple queues
func NewConsumer(rabbitmqURL string, queueNames []string, logger *logger.CustomLogger, es *elasticsearch.Client) (*Consumer, error) {
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

	return &Consumer{
		connection: conn,
		channel:    ch,
		queues:     queueNames,
		log:        logger,
		es:         es,
	}, nil
}

// ConsumeMessages starts consuming messages from all declared queues
func (c *Consumer) ConsumeMessages() {
	for _, queueName := range c.queues {
		go c.consumeQueue(queueName)
	}

	// Block forever
	select {}
}

func (c *Consumer) consumeQueue(queueName string) {
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
		c.log.Errorc(context.TODO(), "Failed to register a consumer for queue", map[string]interface{}{
			"queueName": queueName,
			"error":     err.Error(),
		})
		return
		// log.Fatalf("Failed to register a consumer for queue %s: %s", queueName, err)
	}
	c.log.Infoc(context.TODO(), fmt.Sprintf("Waiting for messages in %s. To exit press CTRL+C", queueName), map[string]interface{}{
		"queueName": queueName,
	})
	// log.Printf("Waiting for messages in %s. To exit press CTRL+C", queueName)
	switch queueName {
	case "code_execution_stats":
		for msg := range msgs {
			err := storeCodeExecutionStatsES(context.TODO(), msg.Body, c.es)
			if err != nil {
				c.log.Errorc(context.TODO(), "error while storing msg in ES", map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			c.log.Infoc(context.TODO(), fmt.Sprintf("Received message from %s: %s", queueName, msg.Body), map[string]interface{}{
				"queueName": queueName,
			})
			// log.Printf("Received message from %s: %s", queueName, msg.Body)
		}
	case "challenge_data":
		for msg := range msgs {
			c.log.Infoc(context.TODO(), fmt.Sprintf("Received message from %s: %s", queueName, msg.Body), map[string]interface{}{
				"queueName": queueName,
			})
		}
	case "user_performance":
		for msg := range msgs {
			c.log.Infoc(context.TODO(), fmt.Sprintf("Received message from %s: %s", queueName, msg.Body), map[string]interface{}{
				"queueName": queueName,
			})
		}

	}
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

func storePerformanceDataES(ctx context.Context, performanceData []byte, es *elasticsearch.Client) error {
	req := esapi.IndexRequest{
		Index:   "user_performance",
		Body:    bytes.NewReader(performanceData),
		Refresh: "true", // To make the document searchable immediately
	}

	resp, err := req.Do(ctx, es)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error storing performance data: %s", resp.String())
	}

	// fmt.Printf("Stored performance data for user %s\n", performanceData.UserID)
	return nil
}

func storeChallengeDataES(ctx context.Context, challengeData []byte, es *elasticsearch.Client) error {
	req := esapi.IndexRequest{
		Index:   "challenge_data",
		Body:    bytes.NewReader(challengeData),
		Refresh: "true", // To make the document searchable immediately
	}

	resp, err := req.Do(ctx, es)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error storing challenge data: %s", resp.String())
	}

	// fmt.Printf("Stored challenge data for challenge %s\n", challengeData.ChallengeID)
	return nil
}
func storeCodeExecutionStatsES(ctx context.Context, codeStats []byte, es *elasticsearch.Client) error {
	req := esapi.IndexRequest{
		Index:   "code_execution_stats",
		Body:    bytes.NewReader(codeStats),
		Refresh: "true", // To make the document searchable immediately
	}

	resp, err := req.Do(ctx, es)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error storing challenge data: %s", resp.String())
	}

	// fmt.Printf("Stored challenge data for challenge %s\n", challengeData.ChallengeID)
	return nil
}
