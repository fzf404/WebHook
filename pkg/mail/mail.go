package mail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

var d *gomail.Dialer
var send []string

// Init Mail
func InitMail(host string, port int, password string, from string, to []string) {
	d = gomail.NewDialer(host, port, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	send = to
}

// Send Mail
func SendMail(subject string, body string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", d.Username)
	m.SetHeader("To", send...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	return d.DialAndSend(m)
}
