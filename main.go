package main

import (
	"github.com/fzf404/WebHook/config"
	"github.com/fzf404/WebHook/log"
	"github.com/fzf404/WebHook/server"
)

func init() {
	log.InitLog()
	config.InitConfig()
	server.InitMail()
	server.InitHook()
}

func main() {
	server.InitServer()
}
