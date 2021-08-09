package main

import (
	"log"
	"webhooks/config"
	"webhooks/shell"

	"net/http"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/webhooks.v5/github"
)

func main() {
	// 初始化Loger文件
	succLoger, errLoger := config.InitLog()

	port := ":" + viper.GetString("port")

	// 批量初始化
	for _, name := range viper.GetStringSlice("list") {

		name := name

		secret := viper.GetString(name + ".secret")
		if secret == "" {
			secret = name
		}

		hookUrl := viper.GetString(name + ".url")
		if hookUrl == "" {
			hookUrl = "/" + name
		}

		shellPath := viper.GetString(name + ".cmd")
		if shellPath == "" {
			shellPath = "./shell/" + name + ".sh"
		}

		secretInit, _ := github.New(github.Options.Secret(secret))
		// 处理函数
		http.HandleFunc(hookUrl, func(w http.ResponseWriter, r *http.Request) {
			// 判断是否为Gitee请求
			if r.Header["User-Agent"][0] == "git-oschina-hook" {

				// 进入secret验证
				log.Print("🚀 In ", name)
				succLoger.Print("🚀 In ", name)

				if r.Header["X-Gitee-Token"][0] != secret {

					log.Print("🚨 Gitee Secret Error.")
					errLoger.Print("🚨 In ", name, ": Gitee Secret Error.")
					return
				}
				go shell.ShellRunner(shellPath, succLoger, errLoger)
				return
			}

			// Github请求处理
			log.Print("🚀 In ", name)
			succLoger.Print("🚀 In ", name)
			payload, err := secretInit.Parse(r, github.PushEvent)
			if err != nil {
				log.Print("🚨 Github Secret Error")
				errLoger.Print("🚨 In ", name, ": Github Secret Error.")
				return
			}
			switch payload := payload.(type) {
			case github.PushPayload:
				// 获得Message
				log.Print("📡 ", payload.HeadCommit.Message)
				succLoger.Print("📡 ", payload.HeadCommit.Message)
				// 执行命令
				go shell.ShellRunner(shellPath, succLoger, errLoger)
			default:
				log.Print("🚨 Undefine Event.")
				errLoger.Print("🚨 In ", name, ": Undefine Event.")
			}

		})
		log.Print("🆕 ", name, ": Init Success.")
	}
	http.ListenAndServe(port, nil)
}
