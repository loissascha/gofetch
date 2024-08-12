package main

import (
	"fmt"
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
		info = strings.Replace(info, "[*user*]", s.username, 1)
	}
	if strings.Contains(info, "[*hostname*]") {
		info = strings.Replace(info, "[*hostname*]", s.hostname, 1)
	}
	if strings.Contains(info, "[*kernelVersion*]") {
		if s.kernelVersion == "" {
			s.loadKernelVersion()
		}
		info = strings.Replace(info, "[*kernelVersion*]", s.kernelVersion, 1)
	}
	if strings.Contains(info, "[*cpuModel*]") {
		info = strings.Replace(info, "[*cpuModel*]", s.cpuModel, 1)
	}
	if strings.Contains(info, "[*desktopSession*]") {
		info = strings.Replace(info, "[*desktopSession*]", s.desktopSession, 1)
	}
	if strings.Contains(info, "[*desktopSessionType*]") {
		info = strings.Replace(info, "[*desktopSessionType*]", s.desktopSessionType, 1)
	}
	if strings.Contains(info, "[*osName*]") {
		info = strings.Replace(info, "[*osName*]", s.osName, 1)
	}
	if strings.Contains(info, "[*uptime*]") {
		info = strings.Replace(info, "[*uptime*]", s.uptime, 1)
	}
	return info
}
