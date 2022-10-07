package log

import (
	"io"
	"log"
	"os"
)

// Init Log
func InitLog() {

	// Init Log Format
	log.SetPrefix("[WebHooks] ")        // 设置日志前缀
	log.SetFlags(log.Ldate | log.Ltime) // 设置日志格式

	// Open Log File
	file, err := os.OpenFile("./log/webhooks.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("🔴 Open Log File Error: \n", err.Error())
	}

	// Set Log Output
	log.SetOutput(io.MultiWriter(file, os.Stdout))
}
