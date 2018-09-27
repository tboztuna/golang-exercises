package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/tboztuna/golang-exercises/url-shortener"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting to server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World!")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
