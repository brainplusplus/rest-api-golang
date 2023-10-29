package services

import (
	"simple-ecommerce/configs"
	"simple-ecommerce/objects"
)

var log = configs.GetLogger()

var (
	smtpCredential *objects.SmtpCredential
	emailFrom      string
)

func InitSmtpCredential() {
	host := configs.GetConfigString("smtp.host")
	port := configs.GetConfigInt("smtp.port")
	username := configs.GetConfigString("smtp.username")
	password := configs.GetConfigString("smtp.password")

	smtpCredential = &objects.SmtpCredential{host, port, username, password}
	emailFrom = configs.GetConfigString("smtp.from_sender")
}
