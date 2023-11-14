package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func main() {
	// Connect to RabbitMQ server
	url := os.Getenv("CLOUDAMQP_URL")
	if url == "" {
		log.Fatal("CLOUDAMQP_URL is not set")
	}

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

	// Declare the queue we will consume from
	q, err := ch.QueueDeclare(
		"test_queue", // queue name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	// Register a consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-acknowledge
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	// Consume messages
	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
			// Here you can handle the message and process it
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
