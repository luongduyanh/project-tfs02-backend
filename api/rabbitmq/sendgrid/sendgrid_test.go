package sendgird

import (
	"fmt"
	"testing"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func TestSendingEmail(t *testing.T) {
	apiKey := "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"
	from := mail.NewEmail("Example Admin", "luongduyanh1999@gmail.com")
	subject := "Thank you for your purchase at shopbase.com"
	to := mail.NewEmail("Example User", "anh.ld172952@sis.hust.edu.vn")
	plainTextContent := "Thank you for purchasing from our store. Here's your order details"
	htmlContent := "<strong>Thank you for purchasing from our store. Here's your order details:</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	t.Error("just call Error function for printing log above")
}
