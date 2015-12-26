package main

import (
	"log"
	"net/http"
)

func main() {
	//Safe as it can only serve files from within the frontend directory
	//At least according to the source but the doc does not mention this
	fileHandler := http.FileServer(http.Dir("./frontend/"))

	http.Handle("/", fileHandler)
	log.Print(http.ListenAndServe(":8080", nil))
}