package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func addFilePath() string {
	args := flag.Args()
	if len(args) == 0 {
		return ""
	}
	return args[0]
}

var (
	countBytesFlag = flag.Bool("c", false, "count the number of bytes")
	countLinesFlag = flag.Bool("l", false, "count the number of lines")
	countWordsFlag = flag.Bool("w", false, "count the number of words")
	countCharsFlag = flag.Bool("m", false, "count the number of characters")
)

func countBytes(scanner *bufio.Scanner) int {
	bytes := 0
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		bytes++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return bytes
}

func countLines(scanner *bufio.Scanner) int {
	lines := 0
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func countWords(scanner *bufio.Scanner) int {
	words := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return words
}

func countCharacters(scanner *bufio.Scanner) int {
	characters := 0
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		characters++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return characters
}

func countAll(scanner *bufio.Scanner) (int, int, int) {
	bytes := 0
	lines := 0
	words := 0

	for scanner.Scan() {
		bytes += len(scanner.Bytes())
		lines++
		words += countWordsInLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return bytes, lines, words
}

func countWordsInLine(line string) int {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func main() {
	flag.Parse()
	filePath := addFilePath()

	var scanner *bufio.Scanner
	if filePath == "" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

	if *countBytesFlag {
		byteCount := countBytes(scanner)
		fmt.Printf("%d %s\n", byteCount, filePath)
	}

	if *countLinesFlag {
		lineCount := countLines(scanner)
		fmt.Printf("%d %s\n", lineCount, filePath)
	}

	if *countWordsFlag {
		wordCount := countWords(scanner)
		fmt.Printf("%d %s\n", wordCount, filePath)
	}

	if *countCharsFlag {
		charCount := countCharacters(scanner)
		fmt.Printf("%d %s\n", charCount, filePath)
	}

	flagHasBeenSet := false
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			flagHasBeenSet = true
		}
	})
	if !flagHasBeenSet {
		byteCount, lineCount, wordCount := countAll(scanner)
		fmt.Printf("%d %d %d %s\n", byteCount, lineCount, wordCount, filePath)
	}
}
