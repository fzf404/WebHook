package server

import (
	"bufio"
	"log"
	"os"
	"strings"

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
			logs, err := tailLines("./log/webhooks.log", 5)
			if err != nil {
				log.Println("无法获取日志文件的最后5行:", err)
			} else {
				logStr := strings.Join(logs, "\n")
				SendMail(name, platform, name+" WebHooks Success", logStr)
			}
		} else {
			// Failure Email
			logs, err := tailLines("./log/webhooks.log", 2)
			if err != nil {
				log.Println("无法获取日志文件的最后2行:", err)
			} else {
				logStr := strings.Join(logs, "\n")
				SendMail(name, platform, name+" WebHooks Failure", logStr)
			}
			//SendMail(name, platform, name+" WebHooks Failure", last)
		}
	}
}

// 从日志文件获取指定行数的最后几行
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
