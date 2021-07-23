package producer

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Producer struct {
	exchange     string
	exchangeType string
	routingKey   string
	channel      *amqp.Channel
}

func CreateNewProducer(exchange, exchangeType, routingKey string, channel *amqp.Channel) *Producer {
	return &Producer{
		exchange:     exchange,
		exchangeType: exchangeType,
		routingKey:   routingKey,
		channel:      channel,
	}
}

// publish message to p.exchange
func (p *Producer) Send(msg string) {
	if p.exchange == "" || p.exchangeType == "" || p.channel == nil {
		fmt.Println("This Producer has a faulty configuration")
		return
	}
	p.declare()
	fmt.Printf("Sending to %v : %v\n", p.exchange, msg)
	err := p.publish(msg)
	if err != nil {
		fmt.Println("Publish msg error: ", err)
	}

}

func (p *Producer) publish(msg string) error {
	err := p.channel.Publish(
		p.exchange,
		p.routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plane",
			ContentEncoding: "",
			Body:            []byte(msg),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) declare() error {
	err := p.channel.ExchangeDeclare(
		p.exchange,     // name
		p.exchangeType, //type
		true,           //durable
		false,          //autoDelete: delete when complete
		false,          //internal
		false,          //noWait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("exchange declare error: %s", err)
	}
	return nil
}

func (p *Producer) Close() error {
	return p.channel.Close()
}
