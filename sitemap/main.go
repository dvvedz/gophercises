package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sitemap-app/linkparser"
)

func getHtml(d string) io.Reader {
	res, err := http.Get(d)

	if err != nil {
		log.Fatal(err)
	}

	return res.Body
}

func redirectHandler(url string) {
	// Where does the url redirect

}

func main() {
	body := getHtml("https://google.com")
	links, err := linkparser.Parse(body)

	if err != nil {
		panic(err)
	}

	for i, l := range links {
		fmt.Println(i, l.Href)
		redirectHandler(l.Href)
	}
}
