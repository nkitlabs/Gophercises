package main

import (
	"flag"
	"fmt"
	"gophercises/cyoa/data"
	"gophercises/cyoa/web"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	fileName := flag.String("file", "data/gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("file not found : %v\n", err)
		return
	}

	story, err := data.ReadStory(file)
	if err != nil {
		fmt.Printf("cannot convert data : %v\n", err)
	}

	h, err := web.NewHandler(story, web.WithTemplateFile("web/web.html"))
	if err != nil {
		log.Printf("canno load handler : %v\n", err)
		return
	}
	portStr := ":" + strconv.Itoa(*port)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting the server on port %s\n", portStr)
	log.Fatal(http.ListenAndServe(portStr, mux))
}
