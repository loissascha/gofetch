package systeminfo

import (
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

func (s *SystemInfo) FillInfoString(info string) string {
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
	if strings.Contains(info, "[*gpuModel*]") {
		if s.gpuModel == "" {
			s.loadGpuModel()
		}
		info = strings.Replace(info, "[*gpuModel*]", s.gpuModel, 1)
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
