package util

import (
	"io"
	"os"
	"os/exec"
)

func Run(command string, r io.Reader, w io.Writer) error {
	var cm *exec.Cmd
	cm = exec.Command("sh", "-c", command)
	cm.Stderr = os.Stderr
	cm.Stdin = r
	cm.Stdout = w
	return cm.Run()
}
