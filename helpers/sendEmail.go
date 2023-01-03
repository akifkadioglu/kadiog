package helpers

import (
	"setup/environment"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(To, Subject, Body string) error {

	mail := gomail.NewMessage()

	mail.SetHeader("From", GoDotEnvVariable(environment.MAIL_FROM_ADDRESS))
	mail.SetHeader("To", To)
	mail.SetHeader("Subject", Subject)
	mail.SetBody("text/html", Body)
	mailPort, _ := strconv.Atoi(GoDotEnvVariable(environment.MAIL_PORT))
	send := gomail.NewDialer(GoDotEnvVariable(environment.MAIL_HOST), mailPort, GoDotEnvVariable(environment.MAIL_FROM_ADDRESS), GoDotEnvVariable(environment.MAIL_PASSWORD))

	if err := send.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
