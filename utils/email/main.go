package Email

import (
	"evermos/config"

	"crypto/tls"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

func Send(to string, subject string, message string) {
	d := gomail.NewDialer(config.MAIL.HOST, 587, config.MAIL.USERNAME, config.MAIL.PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.MAIL.USERNAME)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	err := d.DialAndSend(mailer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("An email has been sent from ", config.MAIL.FROM, " to ", to, " at ", time.Now())
}
