package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	lineNumber   = flag.Bool("n", false, "print the line number")
	noBlankLines = flag.Bool("b", false, "do not print blank lines")
)

func addFilePath() []string {
	args := flag.Args()
	if len(args) == 0 {
		return nil
	} else {
		return args
	}
}

func catOut(scanner bufio.Scanner, line *int) {
	if *lineNumber {
		scanner.Split(bufio.ScanLines)
	}
	for scanner.Scan() {
		text := scanner.Text()
		if *noBlankLines {
			if text == "" {
                fmt.Println(scanner.Text())
			} else {
				*line = *line + 1
				fmt.Printf("%d %s\n", *line, text)
			}
		}
		if *lineNumber {
			*line = *line + 1
			fmt.Printf("%d %s\n", *line, text)

		} else if !*noBlankLines {
			fmt.Println(scanner.Text())
		}
	}
}

func main() {
	flag.Parse()
	filePaths := addFilePath()
	line := 0
	var scanner *bufio.Scanner
	if filePaths == nil {
		scanner = bufio.NewScanner(os.Stdin)
		catOut(*scanner, &line)
	} else {
		for _, filePath := range filePaths {
			f, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			scanner = bufio.NewScanner(f)
			catOut(*scanner, &line)
		}
	}
}
