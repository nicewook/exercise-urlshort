package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercise/urlshort",
		"yaml-godoc":      "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
	- path: /urlshort
		url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
		url: https://github.com/gophercises/urlshort/tree/solution
	`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandlerFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
