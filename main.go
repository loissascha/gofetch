package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

var kernelVersion string
var cpuModel string
var hostname string
var username string
var session string
var osName string
var uptime string

var artStr []string
var longestArtLine int

// TODO:
// - check if config files/folders exist and automatically create them if they don't
// - check if art and info files exist and automatically create them if they don't (one pre defined art and info file by default)
// - randomize art and info files
// - add cli option to define which art and info files to use

func main() {
	initConfig()

	art := getArt()
	infos := getInfo()

	kernelVersion = getKernelVersion()
	cpuModel = getCpuModel()
	hostname = getHostname()
	username = getUsername()
	session = getDesktopSession()
	osName = getOsName()
	uptime = getUptime()

	// find longest art line
	longestArtLine = 0
	for _, v := range art {
		l := removeFormattingFromString(v)
		ln := utf8.RuneCountInString(l)
		if ln > longestArtLine {
			longestArtLine = ln
		}
	}
	longestArtLine += 2

	artStr = []string{}
	for _, v := range art {
		l := v
		for i := utf8.RuneCountInString(removeFormattingFromString(v)); i < longestArtLine; i++ {
			l += " "
		}
		artStr = append(artStr, l)
	}

	maxIndex := len(infos)
	if len(artStr) > maxIndex {
		maxIndex = len(artStr)
	}

	for i := 0; i < maxIndex; i++ {
		fmt.Print("\033[0m") // clear formatting

		fmt.Print(getArtLine(i))

		// if infos are over, skip to next line
		if i >= len(infos) {
			fmt.Print("\033[0m") // clear formatting
			fmt.Print("\n")
			continue
		}

		fmt.Print("\033[0m")
		info := infos[i]
		info = fillInfoString(info)
		fmt.Print(info)
		fmt.Print("\033[0m") // clear formatting
		fmt.Print("\n")
	}
}

func getArtLine(i int) string {
	l := ""
	if i < len(artStr) {
		o := artStr[i]
		if strings.Contains(o, "\\033") {
			o = strings.Replace(o, "\\033", "\033", 100)
		}
		l = o
	} else {
		empty := ""
		for u := 0; u < longestArtLine; u++ {
			empty += " "
		}
		l = empty
	}
	return l
}

func removeFormattingFromString(s string) string {
	if strings.Contains(s, "\\033") {
		s = strings.Replace(s, "\\033", "\033", 100)
	}
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(s, "")
}

func fillInfoString(info string) string {
	if strings.Contains(info, "\\033") {
		info = strings.Replace(info, "\\033", "\033", 100)
	}
	if strings.Contains(info, "[*user*]") {
		info = strings.Replace(info, "[*user*]", username, 1)
	}
	if strings.Contains(info, "[*hostname*]") {
		info = strings.Replace(info, "[*hostname*]", hostname, 1)
	}
	if strings.Contains(info, "[*kernelVersion*]") {
		info = strings.Replace(info, "[*kernelVersion*]", kernelVersion, 1)
	}
	if strings.Contains(info, "[*cpuModel*]") {
		info = strings.Replace(info, "[*cpuModel*]", cpuModel, 1)
	}
	if strings.Contains(info, "[*desktopSession*]") {
		info = strings.Replace(info, "[*desktopSession*]", session, 1)
	}
	if strings.Contains(info, "[*osName*]") {
		info = strings.Replace(info, "[*osName*]", osName, 1)
	}
	if strings.Contains(info, "[*uptime*]") {
		info = strings.Replace(info, "[*uptime*]", uptime, 1)
	}
	return info
}
