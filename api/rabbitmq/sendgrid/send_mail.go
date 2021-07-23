package sendgird

import "fmt"

const ApiKey = "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"

var Sender = EmailUser{
	Name:  "duyanh",
	Email: "luongduyanh1999@gmail.com",
}

func SendVerify(receiverName, receiverEmail string) {
	var receiver = EmailUser{
		Name:  receiverName,
		Email: receiverEmail,
	}

	var emailContent = EmailContent{
		ID:               0,
		Subject:          "Verify order",
		FromUser:         &Sender,
		ToUser:           &receiver,
		PlainTextContent: "_",
		HtmlContent:      "ORDER",
	}

	var sengrid = NewSendgrid(ApiKey)
	//send
	sengrid.Send(&emailContent)
	fmt.Println("sent email")
}
