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
