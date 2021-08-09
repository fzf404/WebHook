package utils

import (
	"log"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	} else {
		log.Print("🚨 Create Log File Success: ", path, err.Error())
		log.Fatal("🚨 Can't Read Log Path:", path, err.Error())
		return false
	}
}

func AutoMkdir(path string) {
	exist := PathExists(path)
	if !exist {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal("🚨 Create Log File Failed: ", path, err.Error())
		} else {
			log.Print("🚨 Create Log File Success: ", path, err.Error())
		}

	}
}
