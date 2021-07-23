package consumer

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Consumer struct {
	exchange     string
	exchangeType string
	bindingKey    string
	queue        string
	channel      *amqp.Channel
}

func CreateNewConsumer(exchange, exchangeType, bindingKey, queue string, channel *amqp.Channel) *Consumer {
	return &Consumer{
		exchange:     exchange,
		exchangeType: exchangeType,
		bindingKey:   bindingKey,
		queue:        queue,
		channel:      channel,
	}
}
func (c *Consumer) StartReceiveData(output chan string) {
	//binding c.queue to c.exchange
	c.bind()

	// consuming data
	msgs := c.consum()
	for {
		data := <-msgs
		fmt.Printf("Receive from %v : %v \n", c.queue, string(data.Body))
		output <- string(data.Body)
	}
}

func (c *Consumer) consum() <-chan amqp.Delivery {
	msgs, err := c.channel.Consume(
		c.queue, // name
		"",      // consumerTag
		true,    //autoAck
		false,   //exclusive
		false,   //noLocal
		false,   //noWait
		nil,     //arguments
	)
	if err != nil {
		fmt.Printf("queue consum error: %v", err)
		return nil
	}
	return msgs
}

//chạy để cài đặt cho queue của Consumer này được binding với exchange có tên trong khai báo của Consumer này.
func (c *Consumer) bind() error {
	// declare exchange
	fmt.Printf("Declare exchange: %v\n", c.exchange)
	if err := c.channel.ExchangeDeclare(
		c.exchange,     // name of the exchange
		c.exchangeType, // type
		true,           // durable
		false,          // delete when complete
		false,          // internal
		false,          // noWait
		nil,            // arguments
	); err != nil {
		return fmt.Errorf("exchange declare error: %s", err)
	}

	// declare queue
	fmt.Printf("Declare queue: %v\n", c.queue)
	queue, err := c.channel.QueueDeclare(
		c.queue, // name of the queue
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // noWait
		nil,     // arguments
	)
	if err != nil {
		return fmt.Errorf("queue declare error: %s", err)
	}

	//binding queue to exchange
	fmt.Printf("Binding queue %v to exchange %v\n", c.queue, c.exchange)
	err2 := c.channel.QueueBind(
		queue.Name,
		c.bindingKey,
		c.exchange,
		false,
		nil,
	)

	if err2 != nil {
		return fmt.Errorf("queue bind error: %s", err)
	}
	return nil
}
func (c *Consumer) Close() error {
	return c.channel.Close()
}
