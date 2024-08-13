package systeminfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s *SystemInfo) loadDesktopSessionType() {
	cmd := exec.Command("sh", "-c", "echo $XDG_SESSION_TYPE")
	r, err := cmd.Output()
	if err != nil {
		fmt.Println("can't read desktop session type")
		return
	}
	st := string(r)
	st = strings.TrimSpace(st)
	st = strings.TrimSuffix(st, "\n")
	s.desktopSessionType = st
	return
}

func (s *SystemInfo) loadDesktopSession() {
	cmd := exec.Command("sh", "-c", "echo $XDG_CURRENT_DESKTOP")
	r, err := cmd.Output()
	if err != nil {
		fmt.Println("can't read desktop session")
		return
	}
	st := string(r)
	st = strings.TrimSpace(st)
	st = strings.TrimSuffix(st, "\n")

	// try with alternate var
	if st == "" {
		cmd := exec.Command("sh", "-c", "echo $XDG_SESSION_DESKTOP")
		r, err := cmd.Output()
		if err != nil {
			fmt.Println("can't read desktop session")
			return
		}
		st = string(r)
		st = strings.TrimSpace(st)
		st = strings.TrimSuffix(st, "\n")
	}
	s.desktopSession = st
	return
}
