package helpers

import (
	"setup/env"
	"fmt"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(To, Subject, Body string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", env.GoDotEnvVariable("MAIL_FROM_ADDRESS"))
	mail.SetHeader("To", To)
	mail.SetHeader("Subject", Subject)
	mail.SetBody("text/plain", Body)
	mailPort, _ := strconv.Atoi(env.GoDotEnvVariable("MAIL_PORT"))
	send := gomail.NewDialer(env.GoDotEnvVariable("MAIL_HOST"), mailPort, env.GoDotEnvVariable("MAIL_FROM_ADDRESS"), env.GoDotEnvVariable("MAIL_PASSWORD"))
	err := send.DialAndSend(mail)
	if err != nil {
		fmt.Println(err.Error())
	}
}
