package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func getArt(artFile string) []string {
	fname := ""
	if artFile != "" {
		fname = artFile + ".txt"
	} else {
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
	return result
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
