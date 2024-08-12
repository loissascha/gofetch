package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type SystemInfo struct {
	kernelVersion      string
	cpuModel           string
	gpuModel           string
	hostname           string
	username           string
	desktopSession     string
	desktopSessionType string
	osName             string
	uptime             string
}

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
	s.desktopSession = st
	return
}

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
	return
}

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
	return
}

func (s *SystemInfo) loadCpuModel() {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("CPUInfo file does not exist!", err)
		return
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
			s.cpuModel = value
			return
		}
	}
	return
}

func (s *SystemInfo) loadUptime() {
	execCmd := exec.Command("uptime", "-p")
	execOut, err := execCmd.Output()
	if err != nil {
		fmt.Println("Error getting uptime: ", err)
		return
	}
	u := string(execOut)
	u = strings.TrimSpace(u)
	u = strings.TrimSuffix(u, "\n")
	s.uptime = u
}

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

func (s *SystemInfo) loadKernelVersion() {
	execCmd := exec.Command("uname", "-sr")
	execOut, err := execCmd.Output()
	if err != nil {
		fmt.Println("Error getting kernel version: ", err)
		return
	}
	k := string(execOut)
	k = strings.TrimSpace(k)
	k = strings.TrimSuffix(k, "\n")
	s.kernelVersion = k
}

func (s *SystemInfo) fillInfoString(info string) string {
	if strings.Contains(info, "\\033") {
		info = strings.Replace(info, "\\033", "\033", 100)
	}
	if strings.Contains(info, "[*user*]") {
		if s.username == "" {
			s.loadUsername()
		}
		info = strings.Replace(info, "[*user*]", s.username, 1)
	}
	if strings.Contains(info, "[*hostname*]") {
		if s.hostname == "" {
			s.loadHostname()
		}
		info = strings.Replace(info, "[*hostname*]", s.hostname, 1)
	}
	if strings.Contains(info, "[*kernelVersion*]") {
		if s.kernelVersion == "" {
			s.loadKernelVersion()
		}
		info = strings.Replace(info, "[*kernelVersion*]", s.kernelVersion, 1)
	}
	if strings.Contains(info, "[*cpuModel*]") {
		if s.cpuModel == "" {
			s.loadCpuModel()
		}
		info = strings.Replace(info, "[*cpuModel*]", s.cpuModel, 1)
	}
	if strings.Contains(info, "[*desktopSession*]") {
		if s.desktopSession == "" {
			s.loadDesktopSession()
		}
		info = strings.Replace(info, "[*desktopSession*]", s.desktopSession, 1)
	}
	if strings.Contains(info, "[*desktopSessionType*]") {
		if s.desktopSessionType == "" {
			s.loadDesktopSessionType()
		}
		info = strings.Replace(info, "[*desktopSessionType*]", s.desktopSessionType, 1)
	}
	if strings.Contains(info, "[*osName*]") {
		if s.osName == "" {
			s.loadOsName()
		}
		info = strings.Replace(info, "[*osName*]", s.osName, 1)
	}
	if strings.Contains(info, "[*uptime*]") {
		if s.uptime == "" {
			s.loadUptime()
		}
		info = strings.Replace(info, "[*uptime*]", s.uptime, 1)
	}
	return info
}
