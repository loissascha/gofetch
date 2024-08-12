package systeminfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s *SystemInfo) loadShell() {
	cmd := exec.Command("sh", "-c", "echo $SHELL")
	execOut, err := cmd.Output()
	if err != nil {
		fmt.Println("Error loading user shell")
		return
	}
	r := string(execOut)
	r = strings.TrimSpace(r)
	r = strings.TrimSuffix(r, "\n")
	if strings.Contains(r, "/usr/bin/") {
		_, after, _ := strings.Cut(r, "/usr/bin/")
		r = after
	}
	s.shell = r
}
