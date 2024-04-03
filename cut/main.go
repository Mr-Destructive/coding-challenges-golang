package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	fieldNumber = flag.Int("f", -1, "print the nth column")
	delimiter   = flag.String("d", "	", "delimiter for the output")
	fieldRange  = flag.String("r", "", "print the lines from the given range lines")
)

func addFilePath() []string {
	args := flag.Args()
	if len(args) == 0 {
		return nil
	} else {
		return args
	}
}

func readContents(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func cutOut(fileReader *os.File) []string {
	reader := csv.NewReader(fileReader)
	reader.Comma = []rune(*delimiter)[0]
	lines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	columns := []string{}
	if *fieldNumber > 0 {
		for _, line := range lines {
			columns = append(columns, line[*fieldNumber-1])
		}
	}
	rows := [][]string{}
	for _, line := range lines {
		rows = append(rows, line)
	}
	if *fieldRange != "" {
		rangeList := strings.Split(*fieldRange, ",")
		start, _ := strconv.Atoi(rangeList[0])
		end, _ := strconv.Atoi(rangeList[1])
		for i := range rows {
			col := rows[i][start-1 : end]
			columns = append(columns, strings.Join(col, " "))
		}
	}
	return columns
}

func main() {
	flag.Parse()
	filePaths := addFilePath()
	if filePaths == nil {
		scanner := bufio.NewScanner(os.Stdin)
		lines := []string{}
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		for _, line := range lines {
			fmt.Println(line)
		}
	} else {
		for _, filePath := range filePaths {
			lines := readContents(filePath)
			f, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			lines = cutOut(f)
			defer f.Close()
			for _, line := range lines {
				fmt.Println(line)
			}
		}
	}
}
