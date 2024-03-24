package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

// Separate function for opening SQL connection Pools by passing DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Command line configuration settings using flags
	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	flag.Parse()

	// Custom defined Leveled logging for Informative and Error logs
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	// DATABASE CONNECTION
	/// Declaring Data source name. Database connection string
	dsn := flag.String("dsn", "sting:Somestrongpass12@/snippetbox?parseTime=true", "MySQL data source name")

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// Initializing application struct to use with URL routing. To share the custom defined loggers from main.go file to others.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Server setup and initialization

	/// Setting up the server by modifying http.Server struct to change the ErrorLog to custom defined errorLog log variable
	/// Routes are called by app.routes() from routes.go file
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Server starting at %s", *addr)
	// Starting the server with custom http.Server struct 'srv'
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
