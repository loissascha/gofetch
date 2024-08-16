package systeminfo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (s *SystemInfo) loadCpuModel() {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("CPUInfo file does not exist!", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	modelName := ""
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "Model" {
			modelName = value
		}

		if key == "model name" {
			s.cpuModel = value
			return
		}
	}

	s.cpuModel = modelName
}
