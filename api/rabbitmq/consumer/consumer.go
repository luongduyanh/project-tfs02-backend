package consumer

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	// "hotel/pkg/gmail"
	sendgird "project-tfs02/api/rabbitmq/sendgrid"
	"sync"

	"github.com/streadway/amqp"
)

// SimpleConsumer a simple consumer structure
type SimpleConsumer struct {
	ctx        context.Context
	wg         *sync.WaitGroup
	channel    *amqp.Channel
	queue      string
	exchange   string
	exchType   string
	bindingKey string
	db         *sql.DB
}

// NewSimpleConsumer creates new consumer
func NewSimpleConsumer(ctx context.Context, wg *sync.WaitGroup, chann *amqp.Channel,
	exchange, exchType, bindingKey, queue string, db *sql.DB) *SimpleConsumer {
	return &SimpleConsumer{
		ctx:        ctx,
		wg:         wg,
		channel:    chann,
		exchange:   exchange,
		exchType:   exchType,
		bindingKey: bindingKey,
		queue:      queue,
		db:         db,
	}
}

// Start start consuming data
func (c *SimpleConsumer) Start() {
	if c.channel == nil || c.queue == "" {
		fmt.Println("Wrong consumer config")
		return
	}
	c.declare()

	fmt.Println("Queue is bound to exchange. Consuming data now")
	msgs, err := c.channel.Consume(
		c.queue, // name
		"",      // consumerTag
		false,   // noAck
		false,   // exclusive
		false,   // noLocal
		false,   // noWait
		nil,     // arguments
	)

	if err != nil {
		fmt.Printf("queue consume error: %v\n", err)
		return
	}

	for {
		select {
		case d := <-msgs:
			// send email
			// update db
			// rebuild em
			em := &sendgird.EmailContent{}
			// string -> emailcontent
			err = json.Unmarshal(d.Body, em)
			if err != nil {
				fmt.Println("Error occurred: ", err)
				continue
			}
			// fmt.Println(em.ToUser.Email)
			// gmail.SendEmailBooking(em)
			// doi send email theo gg api
			// err = c.mailer.Send(em)
			if err != nil {
				fmt.Println("Cannot send email due to error: ", err)
				continue
			}
			// update sql data
			_, err = c.db.Exec("UPDATE `orders` SET confirm_email_sent = ? WHERE id = ?", true, em.ID)
			if err != nil {
				fmt.Println("Cannot update confirm_email_sent to true")
			}
			d.Ack(false) // what is ack false?
		case <-c.ctx.Done():
			fmt.Println("Exiting consumer")
			c.wg.Done()
			return
		}
	}
}

// declare exchange and queue, also bind queue to exchange
func (c *SimpleConsumer) declare() error {
	// declare exchange
	fmt.Printf("Binding exchange %v\n", c.exchange)
	if err := c.channel.ExchangeDeclare(
		c.exchange, // name of the exchange
		c.exchType, // type
		true,       // durable
		false,      // delete when complete
		false,      // internal
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return fmt.Errorf("exchange declare error: %s", err)
	}

	// declare queue
	fmt.Printf("Declare queue %v\n", c.queue)
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

	// binding queue
	fmt.Printf("Binding queue %v to exchange %v\n", c.queue, c.exchange)
	if err = c.channel.QueueBind(
		queue.Name,   // name of the queue
		c.bindingKey, // bindingKey
		c.exchange,   // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("queue bind error: %s", err)
	}
	return nil
}

// Close close consumer
func (c *SimpleConsumer) Close() error {
	return c.channel.Close()
}
