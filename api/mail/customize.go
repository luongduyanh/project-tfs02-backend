package mail

import (
	"fmt"
)

//gửi code xác minh khi người dùng chọn quên mật khẩu
func SendCodeVerify(code int, receiverName, receiverEmail string) {
	var receiver = EmailUser{
		Name:  receiverName,
		Email: receiverEmail,
	}

	var emailContent = EmailContent{
		ID:               0,
		Subject:          "Xác thực tài khoản",
		FromUser:         &Sender,
		ToUser:           &receiver,
		PlainTextContent: "_",
		HtmlContent:      fmt.Sprintf("Mã xác thực là %v", code),
	}

	var sengrid = NewSendgrid(ApiKey)
	//send
	sengrid.Send(&emailContent)
	fmt.Println("sent email")
}

//gửi email thông báo đăng kí thành công
func SendNoticeRegisterSuccessful(user_name, user_email string) {
	var receiver = EmailUser{
		Name:  user_name,
		Email: user_email,
	}
	var emailContent = EmailContent{
		ID:               0,
		Subject:          "Đăng nhập thành công",
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
