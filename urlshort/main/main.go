package main

import (
	"fmt"
	"gophercises/urlshort"
	"net/http"
)

var (
	pathsToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	// space should be used in yaml format instead of pressing tab
	yamlString = `
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution
	`
)

func main() {
	mux := defaultMux()

	yamlPathsToUrls, err := urlshort.BuildMapByYaml([]byte(yamlString))
	if err != nil {
		panic(err)
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	yamlHandler := urlshort.MapHandler(yamlPathsToUrls, mapHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
