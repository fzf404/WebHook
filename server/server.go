package server

import (
	"log"
	"net/http"

	"github.com/fzf404/WebHook/config"
)

// Init Http Server
func InitServer() {

	// Start Http Server
	err := http.ListenAndServe(config.Cfg.Host, nil)
	if err != nil {
		log.Fatal("🔴 http serve error: \n", err.Error())
	}

}
