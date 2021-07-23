package mail

import (
	"fmt"
)

//gửi email thông báo đăng kí thành công
func SendNoticeRegisterSuccessful(user_name, user_email string) {
	var receiver = EmailUser{
		Name:  user_name,
		Email: user_email,
	}
	var emailContent = EmailContent{
		ID:               0,
		Subject:          "Đăng ký thành công",
		FromUser:         &Sender,
		ToUser:           &receiver,
		PlainTextContent: "_",
		HtmlContent:      "Bạn đã đăng ký thành công",
	}

	var sengrid = NewSendgrid(ApiKey)
	//send
	sengrid.Send(&emailContent)
	fmt.Println("sent email")
}
