package main

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func main() {
	// Get RabbitMQ connection URL
	url := os.Getenv("CLOUDAMQP_URL")
	if url == "" {
		log.Fatal("CLOUDAMQP_URL is not set")
	}

	// Connect to RabbitMQ server
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare an exchange
	err = ch.ExchangeDeclare(
		"test_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	// Declare a queue
	q, err := ch.QueueDeclare(
		"test_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Bind the queue to the exchange
	err = ch.QueueBind(
		q.Name,
		"tests", // routing key
		"test_exchange",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}

	// Publish a message
	body := "Hello RabbitMQ!"
	err = ch.PublishWithContext(
		context.Background(),
		"test_exchange",
		"tests", // routing key
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
	log.Printf("Message sent: %s", body)
}
