package mail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

var d *gomail.Dialer
var mailFrom string
var mailTo []string

// InitMailService 初始化邮件服务
func InitMailService(host string, port int, user, pass string, to []string) {
	d = gomail.NewDialer(host, port, user, pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	mailFrom = user
	mailTo = to
}

// SendMail 发送邮件
func SendMail(subject, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mailFrom)
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	return d.DialAndSend(m)
}
