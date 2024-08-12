package systeminfo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (s *SystemInfo) loadHostname() {
	file, err := os.Open("/etc/hostname")
	if err != nil {
		fmt.Println("Hostname not found!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, "\n")
		s.hostname = line
		return
	}
}
