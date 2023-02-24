package helpers

import (
	"setup/env"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(To, Subject, Body string) error {

	mail := gomail.NewMessage()

	mail.SetHeader("From", env.Getenv(env.MAIL_FROM_ADDRESS))
	mail.SetHeader("To", To)
	mail.SetHeader("Subject", Subject)
	mail.SetBody("text/html", Body)
	mailPort, _ := strconv.Atoi(env.Getenv(env.MAIL_PORT))
	send := gomail.NewDialer(env.Getenv(env.MAIL_HOST), mailPort, env.Getenv(env.MAIL_FROM_ADDRESS), env.Getenv(env.MAIL_PASSWORD))

	if err := send.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
