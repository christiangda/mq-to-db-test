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
	consumerName  string
	autoACK       bool
	sendACK       bool
)

func main() {
	flag.StringVar(&serverAddress, "server-address", "127.0.0.1", "RabbitMQ server IP address")
	flag.IntVar(&serverPort, "server-port", 5672, "RabbitMQ server port")
	flag.StringVar(&username, "username", "guest", "RabbitMQ username")
	flag.StringVar(&password, "password", "guest", "RabbitMQ password")
	flag.StringVar(&queue, "queue", "hello", "RabbitMQ Queue")
	flag.StringVar(&consumerName, "consumerName", "helloConsumer", "RabbitMQ Consumer name")
	flag.BoolVar(&autoACK, "autoACK", false, "RabbitMQ Message acknowledgment")
	flag.BoolVar(&sendACK, "sendACK", true, "Sent the ack signal when retrive a message")

	flag.Parse()

	// Connecting
	log.Printf("Connection to: %s", fmt.Sprintf("amqp://%s:%s@%s:%d/\n", username, password, serverAddress, serverPort))
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, serverAddress, serverPort))
	if err != nil {
		log.Fatalf("Failed to connecto to RabbitMQ: %s\n", err)
	}
	defer conn.Close()
	defer log.Println("Closing connection")

	// Go for Channel
	log.Println("Creating a channel")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ Channel: %s\n", err)
	}
	defer ch.Close()
	defer log.Println("Closing channel")

	// Declaring Queue
	log.Printf("Declaring queue: %s\n", queue)
	q, err := ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ Queue: %s\n", err)
	}

	// Consumming messages
	log.Printf("Consumming messages from queue: %s, as consumer name: %s\n", queue, consumerName)
	msgs, err := ch.Consume(
		q.Name,
		consumerName,
		autoACK,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s\n", err)
	}

	log.Println("Waiting to consume messages..., To exit press CTRL+C")
	// here we are reading fron a "read channel", and the range will still waiting until channel close
	for msg := range msgs {
		log.Printf("----------> Message Received: < %s >", msg.Body)
		if sendACK {
			err := msg.Ack(true)
			if err != nil { // tell to RabbitMQ to delete the message because I got it.
				log.Printf("Failed to send ACK: %s", err)
			}
			log.Println("     ...ACK sent")
		} else {
			log.Println("     ...ACK not sent")
		}
	}
	log.Println("bye bye")
}
