package main

import (
	"fmt"
	"gofetch/systeminfo"
	"regexp"
	"strings"
	"unicode/utf8"
)

var sysInfo systeminfo.SystemInfo

var artStr []string
var longestArtLine int

// TODO:
// - randomize art and info files
// - add cli option to define which art and info files to use

func main() {
	initConfig()

	art := getArt()
	infos := getInfo()
	sysInfo = systeminfo.SystemInfo{}

	// ram, _ := sysInfo.ReadMemInfo()
	// fmt.Println(ram)

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
		info = sysInfo.FillInfoString(info)
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
