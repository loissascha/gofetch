package systeminfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s *SystemInfo) loadUsername() {
	cmd := exec.Command("whoami")
	r, err := cmd.Output()
	if err != nil {
		fmt.Println("woami does not exist")
		return
	}
	user := string(r)
	user = strings.TrimSpace(user)
	user = strings.TrimSuffix(user, "\n")
	s.username = user
}
