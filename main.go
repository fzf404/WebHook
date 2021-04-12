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

		secretInit, _ := github.New(github.Options.Secret(secret))
		// 定义处理函数
		http.HandleFunc(hookUrl, func(w http.ResponseWriter, r *http.Request) {
			payload, err := secretInit.Parse(r, github.PushEvent)
			if err != nil {
				if err == github.ErrEventNotFound {
					log.Print(undefineLog)
					return
				}
			}
			log.Print(pushLog)
			switch payload := payload.(type) {
			case github.PushPayload:
				// 获得Message
				log.Print(payload.HeadCommit.Message)
				// 执行命令
				cmd := exec.Command(runCmd)
				stdout, _ := cmd.StdoutPipe()
				cmd.Start()
				bytes, _ := ioutil.ReadAll(stdout)
				log.Print("Run: ", string(bytes))
			}
		})
		log.Print(name, ": 初始化完成")
	}
	http.ListenAndServe(":3000", nil)
}
