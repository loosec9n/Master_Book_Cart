package utils

import (
	"Book_Cart_Project/models"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func SendEmail(user models.User, data *EmailData) error {
	m := gomail.NewMessage()

	//set email sender
	m.SetHeader("From", "bookcartproject@gmail.com")

	//set reciever email address
	m.SetHeader("To", user.Email)

	//set email subject
	m.SetHeader("Subject", data.Subject)

	//ser email body
	m.SetBody("text/plain", data.URL)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "bookcartproject@gmail.com", "justin!23")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
