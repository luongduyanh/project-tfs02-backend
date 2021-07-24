package main

import (
	"project-tfs02/api"
	"project-tfs02/api/mail"
)

func main() {
	go mail.SendNoticeRegisterSuccessful()
	api.Run()
}
