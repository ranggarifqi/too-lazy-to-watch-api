package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func GetRabbitMQClient() *amqp091.Connection {
	url := fmt.Sprintf("amqp://rangga:%s@%s/", os.Getenv("RABBITMQ_DEFAULT_PASS"), os.Getenv("RABBITMQ_HOST"))
	client, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ client: %s", err)
	}
	return client
}

func GetRabbitMQChannel() *amqp091.Channel {
	client := GetRabbitMQClient()
	channel, err := client.Channel()
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ channel: %s", err)
	}
	return channel
}
