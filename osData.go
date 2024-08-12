package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getGpuModel() string {
	cmd := exec.Command("sh", "-c", "lspci | grep -i vga | awk -F ': ' '{print $2}'")
	r, err := cmd.Output()
	if err != nil {
		fmt.Println("can't read gpu model")
		return ""
	}
	s := string(r)
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, "\n")
	return s
}
