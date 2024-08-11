package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInfo() []string {
	file, err := os.Open(configDirPath + "/info/minimal.txt")
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
