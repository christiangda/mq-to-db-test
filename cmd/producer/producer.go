package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var (
	serverAddress string
	serverPort    int
	username      string
	password      string
	queue         string
	message       string
	durable       bool
)

func main() {
	flag.StringVar(&serverAddress, "server-address", "127.0.0.1", "RabbitMQ server IP address")
	flag.IntVar(&serverPort, "server-port", 5672, "RabbitMQ server port")
	flag.StringVar(&username, "username", "guest", "RabbitMQ username")
	flag.StringVar(&password, "password", "guest", "RabbitMQ password")
	flag.StringVar(&queue, "queue", "hello", "RabbitMQ Queue")
	flag.StringVar(&message, "message", "Hello Fucked World!", "Message to send into the queue")
	flag.BoolVar(&durable, "durable", true, "RabbitMQ Message durability")

	flag.Parse()

	// Connecting
	log.Printf("Connection to: %s", fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, serverAddress, serverPort))
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, serverAddress, serverPort))
	if err != nil {
		log.Fatalf("Failed to connecto to RabbitMQ: %s", err)
	}
	defer conn.Close()
	defer log.Println("Closing connection")

	// Go for Channel
	log.Println("Creating a channel")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ Channel: %s", err)
	}
	defer ch.Close()
	defer log.Println("Closing channel")

	// Declaring Queue
	log.Printf("Declaring queue: %s", queue)
	q, err := ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ Queue: %s", err)
	}

	// Publishing
	log.Printf("Publishing message: %s, into the queue: %s", message, queue)
	msgConf := amqp.Publishing{}
	msgConf.ContentType = "text/plain"
	msgConf.Body = []byte(message)
	if durable {
		msgConf.DeliveryMode = amqp.Persistent
	}

	if err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		msgConf,
	); err != nil {
		log.Fatalf("Failed to publish a message on RabbitMQ Queue: %s", err)
	}

	log.Println("Message sent")
	log.Println("bye bye...")
}
