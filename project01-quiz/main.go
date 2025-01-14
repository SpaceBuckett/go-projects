package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var correctAnswers int
	correctAnswers = 0

	csvFileName := flag.String("csv", "problems.csv", "a csv file containing records of questions and answers")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open file %s: Error %s", *csvFileName, err)
		os.Exit(1)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error while reading file: ", err)
	}

	parsedRecords := parseRecords(records)

	for i, record := range parsedRecords {
		var answer string
		fmt.Printf("[%d] Enter Answer for %s\n", i+1, record.question)
		fmt.Scanf("%s\n", &answer)
		if answer == record.answer {
			correctAnswers += 1
		}
	}

	fmt.Printf("🎉 SCORE %d of %d\n", correctAnswers, len(parsedRecords))

}

type problem struct {
	question string
	answer   string
}

func parseRecords(records [][]string) []problem {
	parsedRecords := make([]problem, len(records))

	for i, record := range records {
		parsedRecords[i] = problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
	}
	return parsedRecords
}
