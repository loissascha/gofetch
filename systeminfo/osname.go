package systeminfo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (s *SystemInfo) loadOsName() {
	config := s.readOsRelease()
	on := ""
	for k, v := range config {
		if k == "PRETTY_NAME" {
			on = v
		}
	}
	s.osName = on
}

func (s *SystemInfo) readOsRelease() map[string]string {
	config := map[string]string{}
	file, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return config
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			fmt.Println("Invalid line: ", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"`)

		config[key] = value
	}
	return config
}
