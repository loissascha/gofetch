package main

import (
	"fmt"
	"golang.org/x/term"
)

func getConsoleSize() (int, int) {
	w, h, err := term.GetSize(0)
	if err != nil {
		fmt.Println("error", err)
	}
	return w, h
}
