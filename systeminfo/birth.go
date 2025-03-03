package systeminfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s *SystemInfo) loadBirth() {
	execCmd := exec.Command("stat", "/etc")
	execOut, err := execCmd.Output()
	if err != nil {
		fmt.Println("Error getting birth:", err)
		return
	}
	n := string(execOut)
	split := strings.Split(n, "\n")
	for _, v := range split {
		if strings.Contains(v, "Birth:") {
			b := strings.TrimSpace(v)
			b = strings.TrimPrefix(b, "Birth:")
			b = strings.TrimSpace(b)
			bd := strings.Split(b, " ")
			bt := strings.Split(bd[1], ".")[0]
			s.birth = fmt.Sprintf("%s %s", bd[0], bt)
		}

	}
}
