package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
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

	fmt.Println(parseRecords(records))

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
			answer:   record[1],
		}
	}
	return parsedRecords
}
