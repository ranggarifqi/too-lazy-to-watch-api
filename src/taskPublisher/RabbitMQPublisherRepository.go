package taskPublisher

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type rabbitMQTPublisherRepository struct {
	channel *amqp091.Channel
}

func (r *rabbitMQTPublisherRepository) Publish(channel string, payload PublishPayload) error {

	q, err := r.channel.QueueDeclare(
		channel, // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: payload.ContentType,
			Body:        payload.Body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func NewRabbitMQPublisherRepository(channel *amqp091.Channel) ITaskPublisherRepository {
	return &rabbitMQTPublisherRepository{
		channel: channel,
	}
}
