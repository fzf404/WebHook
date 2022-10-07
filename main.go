package main

import (
	"github.com/fzf404/WebHooks/config"
	"github.com/fzf404/WebHooks/log"
	"github.com/fzf404/WebHooks/server"
)

func init() {
	log.InitLog()       // 初始化日志
	config.InitConfig() // 初始化配置
	server.InitMail()   // 初始化邮件
	server.InitHook()   // 初始化钩子
}

func main() {
	server.InitServer() // 启动服务
}
