package script

import (
	"os/exec"
	"strings"
)

// Run Script
func Run(script string) (string, string) {
	// Structure Script

	cmd := exec.Command("/bin/sh", "-c", script)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", printLog(string(output), 10)
	} else {
		return printLog(string(output), 5), "" // 打印脚本执行日志最后5行
	}

}

// 打印日志的函数，返回指定行数的日志信息
func printLog(logText string, numLines int) string {
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
