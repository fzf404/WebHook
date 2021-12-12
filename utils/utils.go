package utils

import (
	"log"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func AutoMkdir(path string) {
	exist := PathExists(path)
	if !exist {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal("🚨 Create Log File Failed: ", err.Error())
		} else {
			log.Print("🚨 Create Log File Success")
		}
	}
}
