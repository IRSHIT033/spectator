package bootstrap

import (
	"log"

	"github.com/streadway/amqp"
	"spectator.main/internals/rabbitmq"
)

func NewRabbitMQInstance(config *Config) rabbitmq.MQPublisher {
	conn, err := amqp.Dial(config.RabbitMQURI)
	if err != nil {
		log.Fatal(err)
	}

	publisher, err := rabbitmq.NewRabbitMQPublisher(conn, config.RabbitMQQueueName)
	if err != nil {
		publisher.Close()
		log.Fatal(err)
	}
	log.Println("Connected to RabbitMQ")
	return publisher
}
