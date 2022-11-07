package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Read in a quiz provided via a csv file
// the user should be able to provide their own list of problems via a flag
// Default list should be "problems.csv"

// Keep track of how many answers are correct and not correct
// the next querstion should be provided in any case of right or wrong

func quizTimer(d time.Duration) {
	c := time.NewTimer(d * time.Second)
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

func askQuestion(questions [][]string) string {

	reader := bufio.NewReader(os.Stdin)
	var cc int

	for _, row := range questions {
		q := row[0]
		a := row[1]

		fmt.Println("What is:", q)
		fmt.Print("> ")
		quizTimer()
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

func main() {

	var fileFlag = flag.String("f", "problems.csv", "takes a file as arguemnt")
	flag.Parse()

	csv, err := readCsvFile(*fileFlag)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(askQuestion(csv))
}
