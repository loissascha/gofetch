package systeminfo

import (
	"fmt"
	"strings"
	"sync"
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
	memTotal           uint64
	memFree            uint64
	memUsed            uint64
	systemPackages     string
	flatpakPackages    string
	snaps              string
	shell              string
}

func (s *SystemInfo) LoadAllData() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadOsName()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadUptime()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadMemInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadSystemPackages()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		//s.loadFlatpakPackages()
		s.loadFlatpakPackagesv2()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadSnaps()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadShell()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadKernelVersion()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadCpuModel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadGpuModel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadHostname()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadUsername()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadDesktopSession()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.loadDesktopSessionType()
	}()

	func() {
		wg.Wait()
	}()
}

func (s *SystemInfo) FillInfoString(info string) string {
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
		info = strings.Replace(info, "[*kernelVersion*]", s.kernelVersion, 1)
	}
	if strings.Contains(info, "[*cpuModel*]") {
		if s.cpuModel == "" {
			info = strings.Replace(info, "[*cpuModel*]", "Unknown", 1)
		} else {
			info = strings.Replace(info, "[*cpuModel*]", s.cpuModel, 1)
		}
	}
	if strings.Contains(info, "[*gpuModel*]") {
		if s.gpuModel == "" {
			info = strings.Replace(info, "[*gpuModel*]", "Unknown", 1)
		} else {
			info = strings.Replace(info, "[*gpuModel*]", s.gpuModel, 1)
		}
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
	if strings.Contains(info, "[*memUsed*]") {
		var memfloat float64 = float64(s.memUsed)
		memfloat = memfloat / 1000000
		info = strings.Replace(info, "[*memUsed*]", fmt.Sprintf("%.2f", memfloat), 1)
	}
	if strings.Contains(info, "[*memUsedPercent*]") {
		var usedPercent float64 = (100 / float64(s.memTotal)) * float64(s.memUsed)
		info = strings.Replace(info, "[*memUsedPercent*]", fmt.Sprintf("%.0f", usedPercent), 1)
	}
	if strings.Contains(info, "[*memUsedPercentColored*]") {
		usedPercent := (100 / float64(s.memTotal)) * float64(s.memUsed)
		color := "\033[32m"
		if usedPercent >= 50 {
			color = "\033[33m"
		}
		if usedPercent >= 80 {
			color = "\033[31m"
		}
		info = strings.Replace(info, "[*memUsedPercentColored*]", fmt.Sprintf("%v%.0f%v%v", color, usedPercent, "%", "\033[0m"), 1)
	}
	if strings.Contains(info, "[*memFree*]") {
		var memfloat float64 = float64(s.memFree)
		memfloat = memfloat / 1000000
		info = strings.Replace(info, "[*memFree*]", fmt.Sprintf("%.2f", memfloat), 1)
	}
	if strings.Contains(info, "[*memTotal*]") {
		var memfloat float64 = float64(s.memTotal)
		memfloat = memfloat / 1000000
		info = strings.Replace(info, "[*memTotal*]", fmt.Sprintf("%.2f", memfloat), 1)
	}
	if strings.Contains(info, "[*uptime*]") {
		info = strings.Replace(info, "[*uptime*]", s.uptime, 1)
	}
	if strings.Contains(info, "[*packages*]") {
		info = strings.Replace(info, "[*packages*]", s.systemPackages, 1)
	}
	if strings.Contains(info, "[*flatpakPackages*]") {
		info = strings.Replace(info, "[*flatpakPackages*]", s.flatpakPackages, 1)
	}
	if strings.Contains(info, "[*snaps*]") {
		info = strings.Replace(info, "[*snaps*]", s.snaps, 1)
	}
	if strings.Contains(info, "[*shell*]") {
		info = strings.Replace(info, "[*shell*]", s.shell, 1)
	}
	return info
}
