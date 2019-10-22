package gohelper

import (
	"fmt"
	"os/exec"
)

func ExecutorRun(cmd string, args string) (string, error) {
	c := exec.Command("bash", "-c", fmt.Sprintf("%s %s", cmd, args))
	stdout, e := c.CombinedOutput()
	if e != nil {
		return "", e
	}
	return string(stdout), nil
}
