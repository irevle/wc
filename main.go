package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func digits(n int) int {
	w := 1
	for n >= 10 {
		n /= 10
		w++
	}
	return w
}

var shortFlags = "lcwm"

func expandFlags(args []string) ([]string, error) {
	out := make([]string, 0, len(args))
	for _, a := range args {
		if len(a) > 1 && a[0] == '-' && a[1] != '-' {
			for _, c := range a[1:] {
				if !strings.ContainsRune(shortFlags, c) {
					return nil, fmt.Errorf("wc: invalid option -- '%c'", c)
				}
				out = append(out, "-"+string(c))
			}
		} else {
			out = append(out, a)
		}
	}
	return out, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	byteShort := flag.Bool("c", false, "print the byte counts")
	byteLong := flag.Bool("bytes", false, "print the byte counts")

	charShort := flag.Bool("m", false, "print the character counts")
	charLong := flag.Bool("chars", false, "print the character counts")

	newlineShort := flag.Bool("l", false, "print the newline counts")
	newlineLong := flag.Bool("lines", false, "print the newline counts")

	wordShort := flag.Bool("w", false, "print the word counts")
	wordLong := flag.Bool("words", false, "print the word counts")

	args, err := expandFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	flag.CommandLine.Parse(args)
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: wc [-lwc] <filename>")
		os.Exit(1)
	}

	filename := flag.Arg(0)
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount, totalWordCount, charCount, runeCount := 0, 0, 0, 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		lineCount++

		currentLine := scanner.Text()
		charCount += len(currentLine) + 1

		runeCount += utf8.RuneCountInString(currentLine) + 1

		words := strings.Fields(currentLine)
		wordCount := len(words)
		totalWordCount += wordCount
	}

	showBytes := *byteShort || *byteLong
	showChars := *charShort || *charLong
	showNewlines := *newlineShort || *newlineLong
	showWords := *wordShort || *wordLong

	if !showBytes && !showChars && !showNewlines && !showWords {
		showNewlines, showWords, showBytes = true, true, true
	}
	w := 1
	if showNewlines {
		w = max(w, digits(lineCount))
	}
	if showWords {
		w = max(w, digits(totalWordCount))
	}
	if showBytes {
		w = max(w, digits(charCount))
	}
	if showChars {
		w = max(w, digits(runeCount))
	}

	if showNewlines {
		fmt.Printf(" %*d", w, lineCount)
	}
	if showWords {
		fmt.Printf(" %*d", w, totalWordCount)
	}
	if showBytes {
		fmt.Printf(" %*d", w, charCount)
	}
	if showChars {
		fmt.Printf(" %*d", w, runeCount)
	}

	if showBytes || showChars || showNewlines || showWords {
		fmt.Printf(" %s\n", filename)
	}
}
