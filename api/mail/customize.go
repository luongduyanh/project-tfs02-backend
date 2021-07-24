package mail

import (
	"fmt"
	"project-tfs02/api/rabbitMQ/consumer"
	"project-tfs02/api/rabbitMQ/rabbitmq"
)

// Gửi email thông báo đăng kí thành công
func SendNoticeRegisterSuccessful() {
	// Khởi tạo rabbitMQ
	rmq := rabbitmq.CreateNewRMQ("amqp://tfs:tfs-ocg@174.138.40.239:5672/#/")

	// Khởi tạo Channel
	cCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel")
		return
	}

	// Khởi tạo consumer
	consumer := consumer.CreateNewConsumer("emailRegister", "direct", "abc", "emailRegisterQueue", cCh)

	//tao channel de nhan du lieu lay ve
	receiverEmail := make(chan string)

	var sengrid = NewSendgrid(ApiKey)
	// Lấy email về rabbitMQ
	go consumer.StartReceiveData(receiverEmail)

	// Gửi email
	var new_email string
	go func() {
		for {
			new_email = <-receiverEmail
			if new_email != "" {
				var receiver = EmailUser{
					Name:  "new user",
					Email: new_email,
				}
				var emailContent = EmailContent{
					ID:               0,
					Subject:          "Đăng kí thành công",
					FromUser:         &Sender,
					ToUser:           &receiver,
					PlainTextContent: "_",
					HtmlContent:      "Tài khoản của bạn đã đăng kí thành công",
				}
				//send
				sengrid.Send(&emailContent)
				fmt.Println("sent email")
			}
		}
	}()

}
