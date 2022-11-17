package main

import (
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func sendRequest(url string) io.Reader {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// body, readErr := io.ReadAll(res.Body)

	// if readErr != nil {
	// panic(readErr)
	// }

	return res.Body
}

func parseHtml(b io.Reader) {
	doc, err := html.Parse(b)

	if err != nil {
		panic(err)
	}
}

func main() {
	// Get html
	body := sendRequest("https://google.com")
	//fmt.Println(string(res))

	// parse a tags
	parseHtml(body)
}
