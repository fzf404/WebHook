package server

import (
	"log"
	"net/http"

	"github.com/fzf404/WebHooks/config"
	"github.com/fzf404/WebHooks/pkg/hook"
)

// Init Hook
func InitHook() {

	// Traverse Hook
	for k, v := range config.Cfg.Hook {

		name := k
		url := v.Url
		secret := v.Secret
		run := v.Run

		// Handle Http Request
		http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {

			// Judge Platform
			platform, err := hook.JudgePlatform(r)
			if err != nil {
				// Return 403
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("platform undefined"))
				log.Printf("🟡 [%s] unknow platform: %s", name, platform)
				return
			}

			// Validate Secret
			if ok := hook.ValidatePlatform(r, platform, secret); !ok {
				// Return 403
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("secret wrong"))
				log.Printf("🟡 [%s] %s wrong secret : %s", name, platform, secret)
				return
			}

			// Judge Event
			event, err := hook.HandlePlatform(r, platform)
			if err != nil {
				// Return 404
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("event undefined"))
				log.Printf("🔵 [%s] %s unknow event: %s", name, platform, event)
				return
			}

			// Judge Script
			if run[event] == "" {
				// Return 404
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("event undefined"))
				log.Printf("🟡 [%s] %s event: %s", name, platform, event)
				return
			}

			// Return 200
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
			log.Printf("🟢 [%s] %s event: %s", name, platform, event)

			// Execute Script
			go ExecuteScript(name, platform, event, run[event])
		})

		log.Printf("🟢 [%s] listen: http://%s", name, config.Cfg.Host+url)
	}
}
