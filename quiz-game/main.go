package main

import (
	"strings"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("file", "questions.csv", "a CSV file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file: %s\n", csvFileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided file.")
	}

	problems := parseLines(lines)

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1])
		} 
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}
