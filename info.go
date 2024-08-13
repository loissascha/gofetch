package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func getInfo(infoFile string) []string {
	fname := ""
	if infoFile != "" {
		fname = infoFile + ".txt"
	} else {
		fname = getRandomInfoFile()
	}
	file, err := os.Open(configDirPath + "/info/" + fname)
	if err != nil {
		fmt.Println("info not found")
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

func getRandomInfoFile() string {
	fs, err := os.ReadDir(configDirPath + "/info")
	if err != nil {
		fmt.Println("info directory not found")
		return ""
	}
	if len(fs) == 0 {
		fmt.Println("info directory is empty")
		return ""
	}
	i := rand.Intn(len(fs))
	return fs[i].Name()
}
