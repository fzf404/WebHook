package config

import (
	"io/ioutil"
	"log"
	"os"
	"webhooks/utils"

	"gopkg.in/yaml.v2"
)

func InitLog() (*log.Logger, *log.Logger, map[string]interface{}) {
	// 建立 map
	configMap := make(map[string]interface{})
	// 读取配置文件
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal("🚨 Open Config Failed: ", err.Error())
	}
	// 解析配置文件
	if yaml.Unmarshal(yamlFile, configMap) != nil {
		log.Fatal("🚨 Read `config.yaml Error: ", err.Error())
	}
	// 读取日志文件位置
	logPath, succ := configMap["log"].(string)
	if !succ {
		log.Fatal("🚨 Read `config.yaml` Error: log")
	}
	succFile := "success.log"
	errFile := "error.log"
	// 自动新建文件夹
	utils.AutoMkdir(logPath)
	// 成功日志
	succLogFile, err := os.OpenFile(logPath+succFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal("🚨 Open Succ Log File Failed: ", logPath+succFile, err.Error())
	}
	// 失败日志
	errLogFile, err := os.OpenFile(logPath+errFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal("🚨 Open Err Log File Failed: ", logPath+succFile, err.Error())
	}

	succLoger := log.New(succLogFile, "", log.LstdFlags|log.Lshortfile|log.LUTC)
	errLoger := log.New(errLogFile, "", log.LstdFlags|log.Lshortfile|log.LUTC)

	return succLoger, errLoger, configMap
}
