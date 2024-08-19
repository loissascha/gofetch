package main

import (
	"bufio"
	"fmt"
	"gofetch/systeminfo"
	"math/rand"
	"os"
)

func getArt(artFile string, width int, longestInfoLine int, sysInfo *systeminfo.SystemInfo) []string {
	fname := ""
	randomFile := true
	if artFile != "" {
		fname = artFile + ".txt"
		randomFile = false
	}
	tries := 0
	for true {
		tries++
		if randomFile {
			fname = getRandomArtFile()
		}
		file, err := os.Open(configDirPath + "/art/" + fname)
		if err != nil {
			fmt.Println("art not found")
			return []string{}
		}
		defer file.Close()

		result := []string{}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			l := string(line)
			result = append(result, l)
		}

		if randomFile {
			longestLine := getLongestLineLength(result, sysInfo)
			if longestLine > (width - longestInfoLine - 2) {
				if tries >= 100 {
					break
				}
				continue
			}
		}
		return result
	}
	return []string{}
}

func getRandomArtFile() string {
	fs, err := os.ReadDir(configDirPath + "/art")
	if err != nil {
		fmt.Println("art directory not found")
		return ""
	}
	if len(fs) == 0 {
		fmt.Println("art directory is empty")
		return ""
	}
	i := rand.Intn(len(fs))
	return fs[i].Name()
}
