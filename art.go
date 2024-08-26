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
	for true {
		if randomFile {
			fname = getRandomArtFile(longestInfoLine, sysInfo, width)
		}
		if fname == "" {
			break
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
		return result
	}
	return []string{}
}

// get random art file name from all art files which fit in line length
func getRandomArtFile(longestInfoLine int, sysInfo *systeminfo.SystemInfo, width int) string {
	artFiles := getArtFilesInLineLength(longestInfoLine, sysInfo, width)
	if len(artFiles) == 0 {
		return ""
	}
	i := rand.Intn(len(artFiles))
	return artFiles[i]
}

// only get art files which would fit in the current console size
func getArtFilesInLineLength(longestInfoLine int, sysInfo *systeminfo.SystemInfo, width int) []string {
	res := []string{}
	fs, err := os.ReadDir(configDirPath + "/art")
	if err != nil {
		fmt.Println("art directory not found")
		return res
	}
	if len(fs) == 0 {
		fmt.Println("art directory is empty")
		return res
	}
	for _, v := range fs {
		if v.IsDir() {
			continue
		}
		file, err := os.Open(configDirPath + "/art/" + v.Name())
		if err != nil {
			fmt.Println("art not found")
			continue
		}
		defer file.Close()

		result := []string{}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			l := string(line)
			result = append(result, l)
		}
		longestLine := getLongestLineLength(result, sysInfo)
		if longestLine > (width - longestInfoLine - 2) {
			continue
		}
		// add to array
		res = append(res, v.Name())
	}
	return res
}
