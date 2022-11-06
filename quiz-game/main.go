package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Read in a quiz provided via a csv file
// the user should be able to provide their own list of problems via a flag
// Default list should be "problems.csv"

// Keep track of how many answers are correct and not correct
// the next querstion should be provided in any case of right or wrong

func readCsvFile(fn string) ([][]string, error) {
	// read file
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

func askQuestion(questions [][]string) string {

	var qs string
	var cc int

	for _, row := range questions {
		q := row[0]
		a := row[1]

		fmt.Println("What is:", q)
		fmt.Print("> ")
		fmt.Scan(&qs)

		if qs == a {
			fmt.Println("Correct")
			cc += 1
		} else {
			fmt.Println("Incorrect")
		}
	}
	return fmt.Sprintf("You got %d answers right out of %d questions\n", cc, len(questions))
}

func main() {
	csv, err := readCsvFile("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(askQuestion(csv))
}
