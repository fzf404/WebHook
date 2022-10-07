package script

import (
	"io/ioutil"
	"os/exec"
)

// Run Script
func Run(script string) (string, string) {
	// Structure Script
	cmd := exec.Command("/bin/sh", "-c", script)

	// Structure Output
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	// Run Script
	err := cmd.Start()
	if err != nil {
		return "", "script start error"
	}

	// Read Output
	succ, _ := ioutil.ReadAll(stdout)
	fail, _ := ioutil.ReadAll(stderr)

	return string(succ), string(fail)
}
