package main

import (
	"io/ioutil"
	"log"
	"os/exec"

	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	blogPath = "/blog"
	notePath = "/note"
	homePath = "/home"
)

func main() {
	blogHook, _ := github.New(github.Options.Secret("blog"))
	noteHook, _ := github.New(github.Options.Secret("note"))
	homeHook, _ := github.New(github.Options.Secret("home"))

	// blog
	http.HandleFunc(blogPath, func(w http.ResponseWriter, r *http.Request) {
		payload, err := blogHook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				log.Print("🚨: Blog Not Push Event")
				return
			}
		}
		log.Print("🚨: In Blog")
		switch payload.(type) {
		case github.PushPayload:
			// 获得Message
			log.Print(payload.(github.PushPayload).HeadCommit.Message)
			// 执行命令
			cmd := exec.Command("/bin/bash", "/opt/webhooks/shell/blog.sh")
			stdout, _ := cmd.StdoutPipe()
			cmd.Start()
			bytes, _ := ioutil.ReadAll(stdout)
			log.Print("Run: ", string(bytes))
		}
	})

	// note
	http.HandleFunc(notePath, func(w http.ResponseWriter, r *http.Request) {
		payload, err := noteHook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				log.Print("🚨: Note Not Push Event")
				return
			}
		}
		log.Print("🚨: In Note")
		switch payload.(type) {
		case github.PushPayload:
			// 获得Message
			log.Print(payload.(github.PushPayload).HeadCommit.Message)
			// 执行命令
			cmd := exec.Command("/bin/bash", "/opt/webhooks/shell/note.sh")
			stdout, _ := cmd.StdoutPipe()
			cmd.Start()
			bytes, _ := ioutil.ReadAll(stdout)
			log.Print("Run: ", string(bytes))
		}
	})

	// home
	http.HandleFunc(homePath, func(w http.ResponseWriter, r *http.Request) {
		payload, err := homeHook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				log.Print("🚨: Home Not Push Event")
				return
			}
		}
		log.Print("🚨: In Home")
		switch payload.(type) {
		case github.PushPayload:
			// 获得Message
			log.Print(payload.(github.PushPayload).HeadCommit.Message)
			// 执行命令
			cmd := exec.Command("/bin/bash", "/opt/webhooks/shell/home.sh")
			stdout, _ := cmd.StdoutPipe()
			cmd.Start()
			bytes, _ := ioutil.ReadAll(stdout)
			log.Print("Run: ", string(bytes))
		}
	})

	http.ListenAndServe(":3000", nil)
}
