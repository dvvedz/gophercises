package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// TODO: rewrite code to use two different functions for reading one question, and reading all question

func quizTimer(d int) {
	c := time.NewTimer(time.Duration(d) * time.Second)
	<-c.C

	if !c.Stop() {
		fmt.Println()
		fmt.Println("To slow try again, increase time with -duration. Default duration is 30 seconds")
		os.Exit(0)
	}
}

func readCsvFile(fn string) ([][]string, error) {
	f, err := os.ReadFile(fn)
	if err != nil {
		return nil, errors.New("could not read file")
	}

	sf := string(f)

	r := csv.NewReader(strings.NewReader(sf))

	records, readErr := r.ReadAll()

	if readErr != nil {
		return nil, errors.New("invalid csv data")
	}

	return records, nil
}

func askQuestion(questions [][]string, duration int) string {
	reader := bufio.NewReader(os.Stdin)
	var cc int

	for _, row := range questions {
		q := row[0]
		a := row[1]

		fmt.Println("What is:", q)
		fmt.Print("> ")

		quizTimer(duration)

		userInp, _ := reader.ReadString('\n')
		userInp = strings.TrimSuffix(userInp, "\n")

		if userInp == a {
			fmt.Println("Correct")
			cc += 1
		} else {
			fmt.Println("Incorrect")
		}
	}
	return fmt.Sprintf("You got %d answers right out of %d questions\n", cc, len(questions))
}

func shuffleQuestion(q [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q), func(i, j int) {
		q[i], q[j] = q[j], q[i]
	})
	return q
}

func main() {
	var fileFlag = flag.String("f", "problems.csv", "takes a file as arguemnt")
	var shuffleFlag = flag.Bool("shuffle", false, "Shuffle csv data")
	var durationFlag = flag.Int("duration", 30, "Quiz game duration in seconds, quit game if it takes longer then the provided duration")
	flag.Parse()

	csv, err := readCsvFile(*fileFlag)

	if err != nil {
		log.Fatal(err)
	}

	if *shuffleFlag {
		shuffelCsv := shuffleQuestion(csv)
		askQuestion(shuffelCsv, *durationFlag)
	} else {
		fmt.Println(askQuestion(csv, *durationFlag))
	}
}
