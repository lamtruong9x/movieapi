package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// Define a config struct to hold all the configuration settings
type config struct {
	port int
	env  string
}

// Define an application struct to hold the dependencies for HTTP handlers, helpers and middleware.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Read in value for "port" and "env"
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger which writes messages to the standard out stream, // prefixed with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Initialize our application
	app := application{
		config: cfg,
		logger: logger,
	}

	// Declare a new servemux
	mux := app.routes()

	// Create server
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	logger.Printf("starting %s server on %s...", cfg.env, srv.Addr)
	logger.Fatal(srv.ListenAndServe())
}
