package queue

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/streadway/amqp"
	structures "user-service-go/config/configStruct"
)

func ConnectionToRabbit(appData *structures.AppData) (*amqp.Connection, error) {
	spew.Dump("amqp://" + appData.RabbitConfig.RabbitLogin + ":" + appData.RabbitConfig.RabbitPassword + "@" + appData.RabbitConfig.RabbitHost + "/")

	conn, err := amqp.Dial("amqp://" + appData.RabbitConfig.RabbitLogin + ":" + appData.RabbitConfig.RabbitPassword + "@" + appData.RabbitConfig.RabbitHost + "/")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func MakeChannel(connection *amqp.Connection) (*amqp.Channel, error) {
	ch, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"user-service-go.store-service.queue.for.stop.list",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"user-service-go.order-service.queue.for.order.status.update",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
