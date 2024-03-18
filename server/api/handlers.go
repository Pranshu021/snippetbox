package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}


	w.Write([]byte("Welcome to Homepage"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Mothod Not Allowed", http.StatusMethodNotAllowed)
		return
	}


	w.Write([]byte("Create a snippet"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display the snippet of id %d", id)

	// w.Write([]byte("View snippet"))
}