package bootstrap

import (
	"spectator.main/internals/mongo"
	"spectator.main/internals/rabbitmq"
)

type Application struct {
	Config   *Config
	Mongo    mongo.Client
	RabbitMQ rabbitmq.MQPublisher
}

func App() Application {
	app := &Application{}
	app.Config = InitConfig()
	app.Mongo = NewMongoDatabase(app.Config)
	app.RabbitMQ = NewRabbitMQInstance(app.Config)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
