package systeminfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s *SystemInfo) loadGpuModel() {
	cmd := exec.Command("sh", "-c", "lspci | grep -i vga | awk -F ': ' '{print $2}'")
	r, err := cmd.Output()
	if err != nil {
		fmt.Println("can't read gpu model")
		return
	}
	st := string(r)
	st = strings.TrimSpace(st)
	st = strings.TrimSuffix(st, "\n")
	s.gpuModel = st
}
