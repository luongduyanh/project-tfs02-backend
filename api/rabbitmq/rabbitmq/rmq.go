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

// func (rmq *RMQ) QuickCreateNewPairProducerAndConsumer(exchangeName, queueName string, ctx context.Context, wg *sync.WaitGroup) (*producer.Producer, *consumer.Consumer, error) {
// 	//trick configuration :v
// 	routingKey := "abc"
// 	bindingKey := routingKey
// 	exchangeType := "direct"

// 	// create 1 channel for producer
// 	pCh, err := rmq.GetChannel()
// 	if err != nil {
// 		fmt.Println("Cannot get channel: ", err)
// 		return &producer.Producer{}, &consumer.Consumer{}, err
// 	}
// 	// create 1 channel for consumer
// 	cCh, err := rmq.GetChannel()
// 	if err != nil {
// 		fmt.Println("Cannot get channel: ", err)
// 		return &producer.Producer{}, &consumer.Consumer{}, err
// 	}

// 	return producer.CreateNewProducer(exchangeName, exchangeType, routingKey, pCh, ctx, wg),
// 		consumer.CreateNewConsumer(exchangeName, exchangeType, bindingKey, queueName, cCh, ctx, wg),
// 		nil
// }
