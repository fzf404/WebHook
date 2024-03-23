package script

import (
	"os/exec"
	"strings"
)

func RunScript(script string) (string, string) {

	// Run Command
	cmd := exec.Command("/bin/sh", "-c", script)
	output, err := cmd.CombinedOutput()
	

	if err != nil {
		return "", PrintLog(string(output), 8)
	}

	return PrintLog(string(output), 4), ""
}

// Print Log Message
func PrintLog(logText string, numLines int) string {

	lines := strings.Split(strings.TrimSpace(logText), "\n")
	numLogs := len(lines)

	var sb strings.Builder

	if numLogs <= numLines {
		for i := 0; i < numLogs; i++ {
			sb.WriteString(lines[i] + "\n")
		}
	} else {
		for i := numLogs - numLines; i < numLogs; i++ {
			sb.WriteString(lines[i] + "\n")
		}
	}

	return sb.String()
}
