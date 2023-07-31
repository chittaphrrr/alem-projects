package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", http.FileServer(http.Dir("ui/static"))))
	mux.HandleFunc("/", home)

	log.Println("listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err)
		return
	}
}
