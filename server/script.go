package server

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/fzf404/WebHook/config"
	"github.com/fzf404/WebHook/pkg/script"
)

// Execute Script
func ExecuteScript(name string, platform string, event string, cmd string) {

	// Run Script
	succ, fail := script.RunScript(cmd)
	if fail != "" {
		log.Printf("🔴 [%s] %s %s: %s", name, platform, event, fail)
	} else {
		log.Printf("🟢 [%s] %s %s: %s", name, platform, event, succ)
	}

	// Send Email
	if config.Cfg.Mail.Enable {
		if fail == "" {
			// Success Email
			logs, err := tailLines("./log/webhook.log", 4)
			if err != nil {
				log.Println("Read Log Error: ", err)
			} else {
				logStr := strings.Join(logs, "\n")
				SendMail(name, platform, name+" WebHook Success", logStr)
			}
		} else {
			// Fail Email
			logs, err := tailLines("./log/webhook.log", 8)
			if err != nil {
				log.Println("Read Log Error: ", err)
			} else {
				logStr := strings.Join(logs, "\n")
				SendMail(name, platform, name+" WebHook Failure", logStr)
			}
		}
	}
}

// Read Log Message
func tailLines(filePath string, numLines int) ([]string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {

		lines = append(lines, scanner.Text())

		if len(lines) > numLines {
			lines = lines[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
