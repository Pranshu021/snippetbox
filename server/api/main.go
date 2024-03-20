package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
}

func main() {
	// Command line configuration settings using flags
	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	flag.Parse()

	// Custom defined Leveled logging for Informative and Error logs
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)


	// Initializing application struct to use with URL routing. To share the custom defined loggers from main.go file to others.
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	
	mux := http.NewServeMux()

	// Fileserver to serve static files (HTML, CSS, JS)
	fileServer := http.FileServer(http.Dir("client/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// URL Routing
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.HandleFunc("/snippet/view", app.snippetView)

	// Server setup and initialization

	/// Setting up the server by modifying http.Server struct to change the ErrorLog to custom defined errorLog log variable
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}
	infoLog.Printf("Server starting at %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
