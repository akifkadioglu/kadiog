package helpers

import (
	"fmt"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(To, Subject, Body string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", GoDotEnvVariable("MAIL_FROM_ADDRESS"))
	mail.SetHeader("To", To)
	mail.SetHeader("Subject", Subject)
	mail.SetBody("text/plain", Body)
	mailPort, _ := strconv.Atoi(GoDotEnvVariable("MAIL_PORT"))
	send := gomail.NewDialer(GoDotEnvVariable("MAIL_HOST"), mailPort, GoDotEnvVariable("MAIL_FROM_ADDRESS"), GoDotEnvVariable("MAIL_PASSWORD"))
	err := send.DialAndSend(mail)
	if err != nil {
		fmt.Println(err.Error())
	}
}
