package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	lineNumber = flag.Int("n", 10, "print the first n lines")
	byteCount  = flag.Int("c", -1, "print the first n bytes")
)

func addFilePath() []string {
	args := flag.Args()
	if len(args) == 0 {
		return nil
	} else {
		return args
	}
}

func headOut(scanner bufio.Scanner, line *int) []string {
	lines := []string{}
	count := 0
	lineCount := *lineNumber
	if *byteCount > 0 {
		scanner.Split(bufio.ScanBytes)
		lineCount = *byteCount
	}
	var text string
	for scanner.Scan() {
		if *byteCount > 0 {
			text += string(scanner.Bytes())
			if len(text) > lineCount {
				text = text[:lineCount]
				break
			}
		} else {
			text = scanner.Text()
			lines = append(lines, text)
			if count == lineCount-1 {
				break
			}
		}
		count++
	}
	if *byteCount > 0 {
		lines = append(lines, text)
	}
	return lines
}

func main() {
	flag.Parse()
	filePaths := addFilePath()
	line := 0
	var scanner *bufio.Scanner
	if filePaths == nil {
		scanner = bufio.NewScanner(os.Stdin)
		lines := headOut(*scanner, &line)
		for _, line := range lines {
			fmt.Println(line)
		}
	} else {
		for _, filePath := range filePaths {
			f, err := os.Open(filePath)
			fmt.Printf("==> %s <==\n", filePath)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			scanner = bufio.NewScanner(f)
			lines := headOut(*scanner, &line)
			for _, line := range lines {
				fmt.Println(line)
			}
			fmt.Println()
		}
	}
}
