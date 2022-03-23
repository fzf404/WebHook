package shell

import (
	"io/ioutil"
	"log"
	"os/exec"
	"webhooks/mail"
	"webhooks/utils"
)

const (
	MAIL_SUBJECT_SUCCESS = "GoWebHooks Script Exec Successful"
	MAIL_SUBJECT_FAILURE = "GoWebHooks Script Exec Failure"
)

// 执行命令
func ShellRunner(shellPath string, succLoger *log.Logger, errLoger *log.Logger, mailEnable bool) {
	// 判断Shell文件是否存在
	if !utils.PathExists(shellPath) {
		errLoger.Print("🚨 Shell Script Not Exist: ", shellPath)
		log.Fatal("🚨 Shell Script Not Exist: ", shellPath)
	}
	// 执行
	cmd := exec.Command("/bin/bash", shellPath)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		errLoger.Print("🚨 Shell Start Error: ", err.Error())
		log.Fatal("🚨 Shell Start Error.")
	}

	// 读输出
	bytes, _ := ioutil.ReadAll(stdout)
	errbytes, _ := ioutil.ReadAll(stderr)

	if len(errbytes) != 0 {
		log.Print("🚨 Shell Run Error.")
		errLoger.Print("🚨 Shell Run Error: ", string(errbytes))
		if mailEnable {
			mail.SendMail(MAIL_SUBJECT_FAILURE, "🚨 Shell Run Error: "+string(errbytes))
		}
		return
	}

	log.Print("👍 Shell Run Success.")
	succLoger.Print("👍 Shell Run Success: ", string(bytes))
	if mailEnable {
		mail.SendMail(MAIL_SUBJECT_SUCCESS, "👍 Shell Run Success: "+string(bytes))
	}
}
