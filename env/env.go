package env

import (
	"log"
	"os"
)

const (
	DB_CONNECTION = "DB_CONNECTION"
	DB_HOST       = "DB_HOST"
	DB_PORT       = "DB_PORT"
	DB_DATABASE   = "DB_DATABASE"
	DB_USERNAME   = "DB_USERNAME"
	DB_PASSWORD   = "DB_PASSWORD"

	MAIL_MAILER       = "MAIL_MAILER"
	MAIL_HOST         = "MAIL_HOST"
	MAIL_PORT         = "MAIL_PORT"
	MAIL_USERNAME     = "MAIL_USERNAME"
	MAIL_PASSWORD     = "MAIL_PASSWORD"
	MAIL_ENCRYPTION   = "MAIL_ENCRYPTION"
	MAIL_FROM_ADDRESS = "MAIL_FROM_ADDRESS"
	MAIL_FROM_NAME    = "MAIL_FROM_NAME"

	HOST    = "HOST"
	PORT    = "PORT"
	APP_KEY = "APP_KEY"
)

func Init() {

	Setenv(DB_CONNECTION, "mysql")
	Setenv(DB_HOST, "127.0.0.1")
	Setenv(DB_PORT, "3306")
	Setenv(DB_DATABASE, "setup")
	Setenv(DB_USERNAME, "root")
	Setenv(DB_PASSWORD, "root12345")

	Setenv(MAIL_MAILER, "smtp")
	Setenv(MAIL_HOST, "smtp.yaanimail.com")
	Setenv(MAIL_PORT, "587")
	Setenv(MAIL_USERNAME, "")
	Setenv(MAIL_PASSWORD, "")
	Setenv(MAIL_ENCRYPTION, "tls")
	Setenv(MAIL_FROM_ADDRESS, "")
	Setenv(MAIL_FROM_NAME, "")

	Setenv(HOST, "")
	Setenv(PORT, "80")
	Setenv(APP_KEY, "key")
}

func Getenv(key string) string {
	return os.Getenv(key)
}

func Setenv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		log.Printf("%s", err.Error())
	}
}
