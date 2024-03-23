package log

import (
	"io"
	"log"
	"os"
)

// Init Log
func InitLog() {

	// Init Log Format
	log.SetPrefix("[WebHook] ")         // Log Prefix
	log.SetFlags(log.Ldate | log.Ltime) // Log Timestamp

	// Open Log File
	file, err := os.OpenFile("./log/webhook.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("🔴 Open Log File Error: \n", err.Error())
	}

	// Set Log Output
	log.SetOutput(io.MultiWriter(file, os.Stdout))
}
