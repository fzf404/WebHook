package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"webhooks/config"

	"net/http"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/webhooks.v5/github"
)

func main() {
	// 初始化配置文件
	config.InitConfig()
	// 批量初始化
	for _, name := range viper.GetStringSlice("list") {

		secret := viper.GetString(name + ".secret")
		hookUrl := viper.GetString(name + ".url")
		runCmd := viper.GetString(name + ".cmd")
		pushLog := viper.GetString(name + ".push")
		undefineLog := viper.GetString(name + ".undefine")
		gitee := viper.GetBool(name + ".gitee")

		secretInit, _ := github.New(github.Options.Secret(secret))
		// 定义处理函数
		http.HandleFunc(hookUrl, func(w http.ResponseWriter, r *http.Request) {
			if gitee {
				log.Print(pushLog)
				go shellRunner(runCmd)
			} else {
				payload, err := secretInit.Parse(r, github.PushEvent)
				if err != nil {
					log.Print("🚨 Secret Error")
					return
				}
				log.Print(pushLog)
				switch payload := payload.(type) {
				case github.PushPayload:
					// 获得Message
					log.Print(payload.HeadCommit.Message)
					// 执行命令
					go shellRunner(runCmd)
				default:
					log.Print(undefineLog)
				}
			}
		})
		log.Print(name, ": 初始化完成")
	}
	http.ListenAndServe(":3000", nil)
}

func shellRunner(runCmd string) {
	cmd := exec.Command("/bin/bash", runCmd)
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatal("🚨Shell脚本执行错误")
	}
	bytes, _ := ioutil.ReadAll(stdout)
	log.Print("Run: ", string(bytes))
}
