package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	csvFilename := flag.String("f", "problems.csv", "file containing quiz problems")
	timeLimit := flag.Int("t", 30, "quiz timer")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	questions, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	challenges := parseQuestions(questions)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0
	input := ""

	for _, challenge := range challenges {
		fmt.Print(challenge.q, "= ")
		fmt.Scan(&input)
		if input == challenge.a {
			score++
		}
	}
	fmt.Printf("Well done! You scored %d out of %d.\n", score, len(challenges))
}

func parseQuestions(questions [][]string) []question {

	questions_array := make([]question, len(questions))
	for i, row := range questions {
		questions_array[i] = question{q: row[0], a: row[1]}
	}
	return questions_array
}

type question struct {
	q string
	a string
}
