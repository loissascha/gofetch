package systeminfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s *SystemInfo) loadUptime() {
	execCmd := exec.Command("uptime", "-p")
	execOut, err := execCmd.Output()
	if err != nil {
		fmt.Println("Error getting uptime: ", err)
		return
	}
	u := string(execOut)
	u = strings.TrimSpace(u)
	u = strings.TrimSuffix(u, "\n")
	s.uptime = u
}
