package taskPublisher

type ITaskPublisherRepository[T any] interface {
	Publish(channel string, payload T) error
}
