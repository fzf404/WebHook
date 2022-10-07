package server

import (
	"log"

	"github.com/fzf404/WebHooks/config"
	"github.com/fzf404/WebHooks/pkg/script"
)

// Execute Script
func ExecuteScript(name string, platform string, event string, cmd string) {

	// Run Script
	succ, fail := script.Run(cmd)
	if fail != "" {
		log.Printf("🔴 [%s] %s %s: %s", name, platform, event, fail)
	} else {
		log.Printf("🟢 [%s] %s %s: %s", name, platform, event, succ)
	}

	// Send Email
	if config.Cfg.Mail.Enable {
		if fail == "" {
			// Success Email
			last, _ := script.Run("tail -n 4 ./log/webhooks.log | tac")
			SendMail(name, platform, "WebHooks Success", last)
		} else {
			// Failure Email
			last, _ := script.Run("tail -n 8 ./log/webhooks.log | tac")
			SendMail(name, platform, "WebHooks Failure", last)
		}
	}
}
