package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getOsName() string {
	config := readOsRelease()
	osName := ""
	for k, v := range config {
		if k == "PRETTY_NAME" {
			osName = v
		}
	}
	return osName
}

func readOsRelease() map[string]string {
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

func getKernelVersion() string {
	execCmd := exec.Command("uname", "-r")
	execOut, err := execCmd.Output()
	if err != nil {
		fmt.Println("Error getting kernel version: ", err)
		return ""
	}
	k := string(execOut)
	k = strings.TrimSpace(k)
	k = strings.TrimSuffix(k, "\n")
	return k
}

func getUptime() string {
	execCmd := exec.Command("uptime", "-p")
	execOut, err := execCmd.Output()
	if err != nil {
		fmt.Println("Error getting uptime: ", err)
		return ""
	}
	u := string(execOut)
	u = strings.TrimSpace(u)
	u = strings.TrimSuffix(u, "\n")
	return u
}

func getCpuModel() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("CPUInfo file does not exist!", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "model name" {
			return value
		}
	}
	return ""
}

func getHostname() string {
	file, err := os.Open("/etc/hostname")
	if err != nil {
		fmt.Println("Hostname not found!")
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, "\n")
		return line
	}
	return ""
}

func getUsername() string {
	cmd := exec.Command("whoami")
	r, err := cmd.Output()
	if err != nil {
		fmt.Println("woami does not exist")
		return ""
	}
	user := string(r)
	user = strings.TrimSpace(user)
	user = strings.TrimSuffix(user, "\n")
	return user
}

func getDesktopSession() string {
	cmd := exec.Command("sh", "-c", "echo $XDG_CURRENT_DESKTOP")
	r, err := cmd.Output()
	if err != nil {
		fmt.Println("can't read desktop session")
		return ""
	}
	s := string(r)
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, "\n")
	return s
}
