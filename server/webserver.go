package server

import (
	"log"
	"net/http"
)

func init() {
	//Safe as it can only serve files from within the frontend directory
	//At least according to the source but the doc does not mention this
	fileHandler := http.FileServer(http.Dir("../frontend/"))
	http.Handle("/api/task", TaskApi{})

	http.Handle("/", fileHandler)
	log.Print(http.ListenAndServe(":8080", nil))
}
