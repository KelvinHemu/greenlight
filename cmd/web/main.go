package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.kelvinhemu.snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// Define an application struct to hold the dependencies for our web application.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	// Parse the command line flags.
	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Create a new logger instance.
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Open a database connection.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// Close the database connection when the program exits.
	defer db.Close()

	// Create a new application instance.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	// Create a new server and pass it the routes() method.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Start the server.
	infoLog.Println("Starting server on http://localhost" + *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB opens a database connection and returns a pointer to the database object.
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
