package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"
	"webhooks/config"
	"webhooks/shell"

	"net/http"
)

var port string
var succ bool

func init() {
	// 初始化Loger文件
	succLoger, errLoger, configMap := config.InitLog()
	// 获得端口
	port,succ = configMap["port"].(string)
	if !succ {  
		log.Fatal("🚨 Read `config.yaml` Error: port")
	}
	// 获得列表
	list,succ := configMap["list"].([]interface{})
	if !succ {  
		log.Fatal("🚨 Read `config.yaml` Error: list")
	}
	// 批量初始化监听
	for _, name := range list{
		name := name.(string)
		secret := name
		hookUrl := "/" + name
		shellPath := "./shell/" + name + ".sh"
		// 是否在下方覆盖配置
		if config, succ := configMap[name].(map[interface{}]interface{}); succ {
			// 密钥
			if tmp, succ := config["secret"].(string); succ {
				secret = tmp
			}
			// 请求路径
			if tmp, succ := config["url"].(string); succ {
				hookUrl = tmp
			}
			// shell 文件路径
			if tmp, succ := config["cmd"].(string); succ {
				shellPath = tmp
			}
		}

		// 处理函数
		http.HandleFunc(hookUrl, func(w http.ResponseWriter, r *http.Request) {
			// 请求处理
			log.Print("🚀 In ", name)
			succLoger.Print("🚀 In ", name)
			// 获得UA
			userAgent := r.Header.Get("User-Agent")
			switch {
			// Github
			case strings.Contains(userAgent, "GitHub-Hookshot"):
				// 密钥验证
				signature := r.Header.Get("X-Hub-Signature")
				mac := hmac.New(sha1.New, []byte(secret))
				payload, _ := ioutil.ReadAll(r.Body)
				_, _ = mac.Write(payload)
				expectedMAC := hex.EncodeToString(mac.Sum(nil))
				if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
					log.Print("🚨 Github Secret Eroror")
					errLoger.Print("🚨 In ", name, ": Github Secret Error.")
					return
				}
				// Event验证
				switch r.Header.Get("X-Github-Event") {
				case "push":
				case "ping":
					log.Print("🍻 Ping")	
					return 
				default:
					log.Print("🚨 Github Method Error")
					errLoger.Print("🚨 In ", name, ": Github Method Error.")
					return
				}
			// Gitee
			case strings.Contains(userAgent, "git-oschina-hook"):
				// 密钥验证
				if r.Header.Get("X-Gitee-Token") != secret {
					log.Print("🚨 Gitee Secret Error.")
					errLoger.Print("🚨 In ", name, ": Gitee Secret Error.")
					return
				}
				// Event 验证
				switch r.Header.Get("X-Gitee-Event") {
				case "Push Hook":
				default:
					log.Print("🚨 Gitee Method Error")
					errLoger.Print("🚨 In ", name, ": Gitee Method Error.")
					return
				}

			}
			// 运行 Shell 脚本
			go shell.ShellRunner(shellPath, succLoger, errLoger)

		})
		// 初始化成功
		log.Print("🆕 ", name, ": Init Success.")
	}

	// 开启服务
}
func main() {
	http.ListenAndServe(port, nil)
}
