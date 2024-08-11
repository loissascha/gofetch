package main

import (
	"bufio"
	"fmt"
	"os"
)

func getArt() []string {
	homeDir, _ := os.UserHomeDir()
	file, err := os.Open(homeDir + "/.config/go-fetch-tool/art/pikachu2.txt")
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
