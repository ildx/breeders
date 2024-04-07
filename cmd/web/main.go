package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/ildx/breeders/models"
)

const port = ":4000"

type application struct {
	config        appConfig
	Models        models.Models
	templateCache map[string]*template.Template
}

type appConfig struct {
	dsn      string
	useCache bool
}

func main() {
	app := &application{
		templateCache: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	db, err := connectToDatabase(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	app.Models = *models.New(db)

	server := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting server on port", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
