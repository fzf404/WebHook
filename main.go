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

		runCmd := viper.GetString(name + ".cmd")
		if runCmd == "" {
			runCmd = "./shell/" + name + ".sh"
		}

		secretInit, _ := github.New(github.Options.Secret(secret))
		// 定义处理函数
		http.HandleFunc(hookUrl, func(w http.ResponseWriter, r *http.Request) {
			// 判断是否为Gitee请求
			if r.Header["User-Agent"][0] == "git-oschina-hook" {
				log.Print("🚨 In ", name)
				if r.Header["X-Gitee-Token"][0] != "fzf" {
					log.Print("🚨 Gitee Secret Error")
					return
				}
				go shellRunner(runCmd)
				return
			}
			// Github请求处理
			log.Print("🚨 In ", name)
			payload, err := secretInit.Parse(r, github.PushEvent)
			if err != nil {
				log.Print("🚨 Github Secret Error")
				return
			}
			switch payload := payload.(type) {
			case github.PushPayload:
				// 获得Message
				log.Print(payload.HeadCommit.Message)
				// 执行命令
				go shellRunner(runCmd)
			default:
				log.Print("🚨 Undefine Event")
			}

		})
		log.Print(name, ": 初始化完成")
	}
	http.ListenAndServe(port, nil)
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
