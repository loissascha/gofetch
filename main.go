package main

import (
	"flag"
	"fmt"
	"gofetch/systeminfo"
	"regexp"
	"strings"
	"sync"
	"unicode/utf8"
)

var sysInfo systeminfo.SystemInfo

var artStr []string

func main() {
	width, _ := getConsoleSize() // TODO check if max width with art and info fits in the current console window

	var artFile string
	var infoFile string

	flag.StringVar(&artFile, "art", "", "Name  of your art file")
	flag.StringVar(&infoFile, "info", "", "Name  of your info file")

	flag.Parse()

	initConfig()
	sysInfo = systeminfo.SystemInfo{}
	sysInfo.LoadAllData()

	infos := getInfo(infoFile)
	longestInfosLine := getLongestLineLength(infos, &sysInfo)
	// if longestInfosLine > width {
	// 	longestInfosLine = 0
	// }
	art := getArt(artFile, width, longestInfosLine, &sysInfo)

	// find longest art line
	minIndentLength := getLongestLineLength(art, &sysInfo)
	minIndentLength += 2

	artStr = []string{}
	for _, v := range art {
		l := v
		for i := utf8.RuneCountInString(removeFormattingFromString(v)); i < minIndentLength; i++ {
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

		fmt.Print(getArtLine(i, minIndentLength))

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

func getArtLine(i int, minLineLength int) string {
	l := ""
	if i < len(artStr) {
		o := artStr[i]
		if strings.Contains(o, "\\033") {
			o = strings.Replace(o, "\\033", "\033", 100)
		}
		l = o
	} else {
		empty := ""
		for u := 0; u < minLineLength; u++ {
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

func getLongestLineLength(str []string, sysInfo *systeminfo.SystemInfo) int {
	longestLine := 0
	var wg sync.WaitGroup
	for _, v := range str {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l := removeFormattingFromString(v)
			l = sysInfo.FillInfoString(l)

			ln := utf8.RuneCountInString(l)
			if ln > longestLine {
				longestLine = ln
			}
		}()
	}
	func() {
		wg.Wait()
	}()
	return longestLine
}
