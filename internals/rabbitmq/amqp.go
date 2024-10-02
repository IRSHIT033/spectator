package rabbitmq

import (
	"github.com/streadway/amqp"
)

type rabbitMQPublisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

func NewRabbitMQPublisher(conn *amqp.Connection, queueName string) (*rabbitMQPublisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitMQPublisher{
		connection: conn,
		channel:    ch,
		queueName:  queueName,
	}, nil

}

func (p *rabbitMQPublisher) Publish(message []byte) error {
	// Declare the queue (ensure it exists)
	_, err := p.channel.QueueDeclare(
		p.queueName, // queue name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}
	// Publish the message to the queue
	err = p.channel.Publish(
		"",          // exchange
		p.queueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *rabbitMQPublisher) Close() {
	p.channel.Close()
	p.connection.Close()
}
