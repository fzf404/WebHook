package server

import (
	"log"

	"github.com/fzf404/WebHook/config"
	"github.com/fzf404/WebHook/pkg/mail"
)

// Init Mail
func InitMail() {

	// Judge Mail Enable
	if config.Cfg.Mail.Enable {
		mail.InitMail(config.Cfg.Mail.Host, config.Cfg.Mail.Port, config.Cfg.Mail.Password, config.Cfg.Mail.From, config.Cfg.Mail.To)
	}

}

// Send Mail
func SendMail(name string, platform string, subject string, body string) {

	// Send Mail
	if err := mail.SendMail(subject, body); err != nil {
		log.Printf("🔴 [%s] %s mail: %s ", name, platform, err.Error())
		return
	}

	log.Printf("🟢 [%s] %s mail: %v", name, platform, config.Cfg.Mail.To)
}
