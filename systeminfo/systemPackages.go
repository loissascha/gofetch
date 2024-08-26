package systeminfo

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (s *SystemInfo) loadSystemPackages() {
	var cmd *exec.Cmd
	pkgManager := "pacman"
	if !commandExists(pkgManager) {
		pkgManager = "rpm"
		if !commandExists(pkgManager) {
			pkgManager = "dpkg"
			if !commandExists(pkgManager) {
				fmt.Println("No package manager found")
				return
			}
		}
	}
	if pkgManager == "dpkg" {
		cmd = exec.Command("sh", "-c", "dpkg -l | grep '^ii' | wc -l")
	} else if pkgManager == "rpm" {
		cmd = exec.Command("sh", "-c", "rpm -qa | wc -l")
	} else if pkgManager == "pacman" {
		cmd = exec.Command("sh", "-c", "pacman -Q | wc -l")
	} else {
		fmt.Println("Error loading system packages")
		return
	}
	execOut, err := cmd.Output()
	if err != nil {
		fmt.Println("Error loading system packages")
		return
	}
	r := string(execOut)
	r = strings.TrimSpace(r)
	r = strings.TrimSuffix(r, "\n")
	s.systemPackages = "(" + r + ") " + pkgManager
}

func (s *SystemInfo) loadFlatpakPackages() {
	if commandExists("flatpak") {
		cmd := exec.Command("sh", "-c", "find /var/lib/flatpak/app -mindepth 1 -maxdepth 1 -type d | wc -l")
		//cmd := exec.Command("sh", "-c", "flatpak list | wc -l")
		execOut, err := cmd.Output()
		if err != nil {
			fmt.Println("Error loading flatpak packages")
			return
		}
		r := string(execOut)
		r = strings.TrimSpace(r)
		r = strings.TrimSuffix(r, "\n")
		s.flatpakPackages = "(" + r + ") " + "flatpaks"
	}
}

func (s *SystemInfo) loadFlatpakPackagesv2() {
	dir, err := os.ReadDir("/var/lib/flatpak/app")
	if err != nil {
		s.flatpakPackages = ""
		return
	}
	count := len(dir)
	r := fmt.Sprintf("%v", count)
	s.flatpakPackages = "(" + r + ") flatpaks"
}

func (s *SystemInfo) loadSnaps() {
	if commandExists("snap") {
		cmd := exec.Command("sh", "-c", "snap list | grep -v '^Name' | wc -l")
		execOut, err := cmd.Output()
		if err != nil {
			fmt.Println("Error loading snaps")
			return
		}
		r := string(execOut)
		r = strings.TrimSpace(r)
		r = strings.TrimSuffix(r, "\n")
		s.snaps = "(" + r + ") " + "snaps"
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
