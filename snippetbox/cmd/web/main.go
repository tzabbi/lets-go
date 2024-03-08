package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", notFoundDir(fileServer)))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/download", downloadFile)

	http.HandleFunc()
	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
