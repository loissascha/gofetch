package systeminfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (s *SystemInfo) loadMemInfo() {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println("Error reading memory")
		return
	}
	defer file.Close()

	memInfo := make(map[string]uint64)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		key := strings.TrimSuffix(fields[0], ":")
		value, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			fmt.Println("Error reading memory")
			return
		}

		memInfo[key] = value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading memory")
		return
	}

	for k, v := range memInfo {
		if k == "MemTotal" {
			s.memTotal = v
		} else if k == "MemAvailable" {
			s.memFree = v
		}
	}
	s.memUsed = s.memTotal - s.memFree
}
