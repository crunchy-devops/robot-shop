package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	// Define RabbitMQ server URL
	amqpServerURL := "amqp://guest:guest@my-rabbitmq-cluster:5672/"

	// Create a new RabbitMQ connection
	conn, err := amqp.Dial(amqpServerURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	log.Println("Successfully connected to RabbitMQ")

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"test_queue", // name of the queue
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	log.Printf("Declared queue %s", q.Name)

	// Publish a message to the queue
	body := "Hello, RabbitMQ!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key (queue name)
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}

	log.Printf("Sent message: %s", body)
}
