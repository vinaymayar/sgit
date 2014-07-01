package utils

import (
	"bytes"
	"os/exec"
)

func Execute(cmdStr string, args ...string) (string, error) {
	cmd := exec.Command(cmdStr, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func ExecuteWithStderr(cmdStr string, args ...string) (string, string, error) {
	cmd := exec.Command(cmdStr, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
