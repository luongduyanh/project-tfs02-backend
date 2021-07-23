package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RMQ struct {
	URI        string `json:"uri"`
	connection *amqp.Connection
}

func (rmq *RMQ) GetChannel() (*amqp.Channel, error) {
	return rmq.connection.Channel()
}

func (rmq *RMQ) Close() {
	rmq.connection.Close()
}

func CreateNewRMQ(uri string) *RMQ {
	con, err := amqp.Dial(uri)
	if err != nil {
		fmt.Println("Cannot connect to RabbitMQ !")
		return nil
	}
	return &RMQ{
		uri,
		con,
	}
}
