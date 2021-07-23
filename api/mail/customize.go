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
		Subject:          "Verify your account",
		FromUser:         &Sender,
		ToUser:           &receiver,
		PlainTextContent: "_",
		HtmlContent:      fmt.Sprintf("Your verification code is: %v", code),
	}

	var sengrid = NewSendgrid(ApiKey)
	//send
	sengrid.Send(&emailContent)
	fmt.Println("sent email")
}

//gửi thông báo xử lý thành công khi người dụng nhập sản phẩm bằng file
func SendNoticeImportSuccessful(admin_name, admin_email string) {
	var receiver = EmailUser{
		Name:  admin_name,
		Email: admin_email,
	}

	var emailContent = EmailContent{
		ID:               0,
		Subject:          "Import Successfully",
		FromUser:         &Sender,
		ToUser:           &receiver,
		PlainTextContent: "_",
		HtmlContent:      "Your file processing is complete",
	}

	var sengrid = NewSendgrid(ApiKey)
	//send
	sengrid.Send(&emailContent)
	fmt.Println("sent email")
}

//gửi báo cáo thống kê
func SendReport(dataStatistic string) {
	var receiver = EmailUser{
		Name:  "ngoc nguyen",
		Email: "nguyendinhhdpv3@gmail.com",
	}

	var emailContent = EmailContent{
		ID:               0,
		Subject:          "Daily report",
		FromUser:         &Sender,
		ToUser:           &receiver,
		PlainTextContent: "_",
		HtmlContent:      dataStatistic,
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
		Subject:          "Import Successfully",
		FromUser:         &Sender,
		ToUser:           &receiver,
		PlainTextContent: "_",
		HtmlContent:      "Your file processing is complete",
	}

	var sengrid = NewSendgrid(ApiKey)
	//send
	sengrid.Send(&emailContent)
	fmt.Println("sent email")
}
