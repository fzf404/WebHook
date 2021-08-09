package config

import (
	"log"
	"os"
	"webhooks/utils"

	"github.com/spf13/viper"
)

var succLoger *log.Logger
var errLoger *log.Logger

func InitConfig() {
	// 设置配置文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// 搜索路径
	viper.AddConfigPath("./config")
	// 自动根据类型来读取配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("🚨 Read Config Error: ", err)
	}
}

func InitLog() (*log.Logger, *log.Logger) {
	InitConfig()

	logPath := viper.GetString("log")
	succFile := "success.log"
	errFile := "error.log"

	utils.AutoMkdir(logPath)

	// 成功打印
	succLogFile, err := os.OpenFile(logPath+succFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal("🚨 Open Succ Log File Failed: ", logPath+succFile, err.Error())
	}
	// 失败打印
	errLogFile, err := os.OpenFile(logPath+errFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal("🚨 Open Err Log File Failed: ", logPath+succFile, err.Error())
	}

	succLoger := log.New(succLogFile, "", log.LstdFlags|log.Lshortfile|log.LUTC)
	errLoger := log.New(errLogFile, "", log.LstdFlags|log.Lshortfile|log.LUTC)

	return succLoger, errLoger
}
