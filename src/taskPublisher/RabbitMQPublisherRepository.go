package taskPublisher

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type rabbitMQTPublisherRepository struct {
	channel *amqp091.Channel
}

func (r *rabbitMQTPublisherRepository) Publish(channel string, payload amqp091.Publishing) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.channel.PublishWithContext(ctx,
		"",      // exchange
		channel, // routing key
		false,   // mandatory
		false,   // immediate
		payload,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewRabbitMQPublisherRepository(channel *amqp091.Channel) ITaskPublisherRepository[amqp091.Publishing] {
	return &rabbitMQTPublisherRepository{
		channel: channel,
	}
}
