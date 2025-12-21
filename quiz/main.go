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

	csvFile := flag.String("f", "problems.csv", "file containing quiz problems")
	timeLimit := flag.Int("t", 30, "quiz timer")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)
	}

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

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			fmt.Printf("You scored %d out of %d.\n", score, len(challenges))
			return
		default:
			if input == challenge.a {
				score++
			}
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
