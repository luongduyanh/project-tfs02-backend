package producer

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"project-tfs02/api/models"
	sendgird "project-tfs02/api/rabbitmq/sendgrid"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

const (
	DefaultFromName  = "My Store Owner"
	DefaultFromEmail = "support@mystore.com"
)

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// SimpleProducer a simple producer structure
type SimpleProducer struct {
	ctx        context.Context
	wg         *sync.WaitGroup
	channel    *amqp.Channel
	exchange   string
	exchType   string
	routingKey string
	db         *sql.DB
}

// NewSimpleProducer creates new producer
func NewSimpleProducer(ctx context.Context, wg *sync.WaitGroup, chann *amqp.Channel,
	exchange, exchType, routingKey string, db *sql.DB) *SimpleProducer {
	return &SimpleProducer{
		ctx:        ctx,
		wg:         wg,
		channel:    chann,
		exchange:   exchange,
		exchType:   exchType,
		routingKey: routingKey,
		db:         db,
	}
}

// Start start generating data
func (p *SimpleProducer) Start() {
	if p.channel == nil || p.exchType == "" || p.exchange == "" {
		fmt.Println("Wrong producer config")
		return
	}
	// declare exchanges
	p.declare()

	// create a ticker
	ticker := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-ticker.C:
			// scan db & send to rmq
			fmt.Printf("Scanning for new order(s) at %v\n", time.Now().Format("2006-Jan-02 15:04:05"))
			resp, err := p.getEmailForSending()
			if err != nil {
				return
			}
			fmt.Printf("Scheduling %v email(s) at %v\n", len(resp), time.Now().Format("2006-Jan-02 15:04:05"))
			for _, em := range resp {
				b, _ := json.Marshal(em)
				err := p.publish(p.exchange, p.routingKey, string(b))
				if err != nil {
					fmt.Println("error when publishing data: ", err)
				}
			}
		case <-p.ctx.Done():
			fmt.Println("Exiting consumer")
			ticker.Stop()
			p.wg.Done()
			return
		}
	}
}

func (p *SimpleProducer) publish(exch, routingKey, body string) error {
	if err := p.channel.Publish(
		exch,       // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
		},
	); err != nil {
		return fmt.Errorf("publish data error: %s", err)
	}
	return nil
}

// declare exchange and queue, also bind queue to exchange
func (p *SimpleProducer) declare() error {
	// declare exchange
	fmt.Printf("Binding exchange %v\n", p.exchange)
	if err := p.channel.ExchangeDeclare(
		p.exchange, // name of the exchange
		p.exchType, // type
		true,       // durable
		false,      // delete when complete
		false,      // internal
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return fmt.Errorf("exchange declare error: %s", err)
	}
	return nil
}

// Close close producer
func (c *SimpleProducer) Close() error {
	return c.channel.Close()
}

// getEmailForSending get email and fill up enough information ready for sending
func (p *SimpleProducer) getEmailForSending() ([]*sendgird.EmailContent, error) {
	resp, err := p.scanFromDB()
	if err != nil {
		return resp, err
	}
	// fill FromUser
	// why we can set FromUser here?
	for _, emailContent := range resp {
		emailContent.FromUser = &sendgird.EmailUser{
			Name:  DefaultFromName,
			Email: DefaultFromEmail,
		}
	}

	return resp, err
}

// scanFromDB get all orders that match the predefined condition (created_at < now - 1 min && thankyou_email_sent == falses)
func (p *SimpleProducer) scanFromDB() ([]*sendgird.EmailContent, error) {
	var resp []*sendgird.EmailContent
	// fromTime := time.Now().Add(-time.Minute * 2) // subtract by 2 minutes - why not one?
	// What is prepared statement? Why we should know and use that? is the below usage right? Why not?
	// stmt, err := p.db.Prepare("SELECT id, customer_name, email FROM `order` WHERE created_at >= ? AND thankyou_email_sent = ?;")
	stmt, err := p.db.Prepare("SELECT id,user_id,status_id,total_price FROM `orders` WHERE confirm_email_sent = ?;")
	if err != nil {
		fmt.Println("Cannot prepare statement, ", err)
		return nil, err
	}
	// rows, err := stmt.Query(fromTime, false)
	rows, err := stmt.Query(false)
	if err != nil || rows == nil {
		fmt.Printf("Cannot query from db due to error: %v, %v\n", err, rows == nil)
		return nil, err
	}
	// MUST to call this function at the end to free connection to mysql
	defer rows.Close()

	var id, user_id, status_id uint
	var total_price string
	for rows.Next() {
		// err = rows.Scan(&id, &name, &email)
		err = rows.Scan(&id, &user_id, &status_id, &total_price)
		if err != nil {
			fmt.Println("Cannot scan row due to error: ", err)
			continue
		}
		// var hotel Hotel
		// err = p.db.QueryRow("SELECT id, name, user_id FROM hotels where id = ?", hotel_id).Scan(&hotel.ID, &hotel.Name, &hotel.UserID)
		// if err != nil {
		// 	panic(err.Error()) // proper error handling instead of panic in your app
		// }
		var user models.User

		err = p.db.QueryRow("SELECT id, name, email FROM users where id = ?", user_id).Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// err = p.db.QueryRow("SELECT id, first_name, last_name, email FROM users where id = ?", user_id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		// if err != nil {
		// 	panic(err.Error()) // proper error handling instead of panic in your app
		// }
		// fmt.Println(Hotelier.Email, user.Email)
		// fmt.Println(id, user_id, hotel_id, room_id, time_id, total)
		resp = append(resp, &sendgird.EmailContent{
			ID:               id,
			Subject:          "Hi " + user.Name + ", thank you for purchase at shopbase.com",
			PlainTextContent: "Hi " + user.Name + ", Thank you for booking from us",
			HtmlContent:      "<strong>Here are your order details:</strong>",
			ToUser: &sendgird.EmailUser{
				Name:  user.Name,
				Email: user.Email,
			},
		})
	}
	return resp, nil
}
