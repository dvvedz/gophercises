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

// TODO add timer

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
	reader := bufio.NewReader(os.Stdin)
	var cc int

	for _, row := range questions {
		q := row[0]
		a := row[1]

		fmt.Println("What is:", q)
		fmt.Print("> ")
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
	flag.Parse()

	csv, err := readCsvFile(*fileFlag)

	if err != nil {
		log.Fatal(err)
	}

	if *shuffleFlag {
		shuffelCsv := shuffleQuestion(csv)
		askQuestion(shuffelCsv)
	} else {
		fmt.Println(askQuestion(csv))
	}
}
