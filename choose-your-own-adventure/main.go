package main

import (
	"flag"
	"fmt"
	"os"
	"textadventure/story"
)

func readFile(file string) []byte {
	val, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	file := flag.String("file", "story.json", "the JSON file with the story")
	flag.Parse()

	fmt.Println("Using the story:", *file)

	jd := readFile(*file)

	getJson, err := story.HandleJson(jd)
	if err != nil {
		panic(err)
	}

	fmt.Println(getJson)
}
