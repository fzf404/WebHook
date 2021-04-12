package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	// 设置配置文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// 搜索路径
	viper.AddConfigPath("./config")
	// 自动根据类型来读取配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}
}
