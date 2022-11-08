package main

import (
	"fmt"
	"net/http"
	"urlshort/urlshort"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Url shortener application") })

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/g":  "https://google.com",
		"/gi": "https://github.com",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, r)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/final
`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", yamlHandler)
}
