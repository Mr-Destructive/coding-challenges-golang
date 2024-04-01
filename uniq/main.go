package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	lineCountFlag = flag.Bool("c", false, "count the number of lines")
	repeatedLines = flag.Bool("d", false, "print the non-unique lines")
	onlyUniqFlag  = flag.Bool("u", false, "print only the uniq lines")
)

func addFilePath() []string {
	args := flag.Args()
	if len(args) == 0 {
		return nil
	} else {
		return args
	}
}

func inMap(list map[int]string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func uniqOut(scanner *bufio.Scanner) []string {
	scanner.Split(bufio.ScanLines)
	lines := make(map[int]string)
	lineCount := make(map[string]int)
	n := 0
	uniqLines := []string{}
	nonUniqLines := []string{}
	linesList := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if !inMap(lines, text) {
			n += 1
			lines[n] = text
			uniqLines = append(uniqLines, text)
		} else {
			nonUniqLines = append(nonUniqLines, text)
		}
		lineCount[text] += 1
	}
	if *repeatedLines {
		linesList = nonUniqLines
		lines := []string{}
		if *lineCountFlag {
			for _, line := range linesList {
				for i := 0; i < lineCount[line]; i++ {
					lines = append(lines, line)
				}
			}
			linesList = lines
		}
	} else {
		linesList = uniqLines
	}
	if *onlyUniqFlag {
		uniqLines := []string{}
		for _, line := range linesList {
			if lineCount[line] == 1 {
				uniqLines = append(uniqLines, line)
			}
		}
		linesList = uniqLines
	}
	if *lineCountFlag {
		linesCount := []string{}
		for _, line := range linesList {
			line = fmt.Sprintf("%d %s", lineCount[line], line)
			linesCount = append(linesCount, line)
		}
		return linesCount
	}
	return linesList
}

func main() {
	flag.Parse()
	filePaths := addFilePath()
	var scanner *bufio.Scanner
	if filePaths == nil {
		scanner = bufio.NewScanner(os.Stdin)
		for _, lines := range uniqOut(scanner) {
			fmt.Println(lines)
		}
	} else {
		for _, filePath := range filePaths {
			f, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			scanner = bufio.NewScanner(f)
			uniqLines := uniqOut(scanner)
			for _, lines := range uniqLines {
				fmt.Println(lines)
			}
		}
	}
}
