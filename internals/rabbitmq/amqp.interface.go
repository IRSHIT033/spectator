package rabbitmq

type MQPublisher interface {
	Publish(message []byte) error
}
