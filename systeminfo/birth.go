package systeminfo

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
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
	layout := "2006-01-02 15:04:05"
	for _, v := range split {
		if strings.Contains(v, "Birth:") {
			birthStr := strings.TrimSpace(v)
			birthStr = strings.TrimPrefix(birthStr, "Birth:")
			birthStr = strings.TrimSpace(birthStr)
			birthDate := strings.Split(birthStr, " ")
			birthTime := strings.Split(birthDate[1], ".")[0]
			timestr := fmt.Sprintf("%s %s", birthDate[0], birthTime)
			parsedTime, err := time.Parse(layout, timestr)
			if err != nil {
				s.birth = "How unfortunate... your time string is malformed."
				return
			}
			now := time.Now()
			duration := now.Sub(parsedTime)
			days := int(duration.Hours() / 24)
			s.birth = fmt.Sprintf("%v days", days)
		}

	}
}
