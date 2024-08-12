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

	st = s.extractGpuModel(st)

	s.gpuModel = st
}

func (s *SystemInfo) extractGpuModel(st string) string {
	runTest := true
	ss := st
	res := ""
	for runTest {
		_, after, foundStart := strings.Cut(ss, "[")
		if foundStart {
			before, end, foundEnd := strings.Cut(after, "]")
			if foundEnd {
				res += before + " "
				ss = end
			} else {
				runTest = false
				break
			}
		} else {
			runTest = false
			break
		}
	}
	if res == "" {
		res = st
	}
	return res
}
