package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("client/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)

	log.Print("[+] Server starting at port 4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
