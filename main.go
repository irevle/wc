package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount, totalWordCount, charCount := 0, 0, 0

	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		lineCount++
		currentLine := scanner.Text()
		charCount += len(currentLine) + 1

		words := strings.Fields(currentLine)
		wordCount := len(words)

		totalWordCount += wordCount
	}
	fmt.Println("---- FINAL OUTPUT ----")
	fmt.Printf("%2v %2v %2v %2v\n", lineCount, totalWordCount, charCount, filename)
}
