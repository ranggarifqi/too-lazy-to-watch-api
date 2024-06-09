package taskPublisher

type PublishPayload struct {
	ContentType string
	Body        []byte
}

type ITaskPublisherRepository interface {
	Publish(channel string, payload PublishPayload) error
}
