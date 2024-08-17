package systeminfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s *SystemInfo) loadKernelVersion() {
	execCmd := exec.Command("uname", "-msr")
	execOut, err := execCmd.Output()
	if err != nil {
		fmt.Println("Error getting kernel version: ", err)
		return
	}
	k := string(execOut)
	k = strings.TrimSpace(k)
	k = strings.TrimSuffix(k, "\n")
	s.kernelVersion = k
}
