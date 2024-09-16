package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/go-chi/chi/v5"
)

func main() {

	// Get DB Vars
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSsl := os.Getenv("DB_SSL")
	backendPort := os.Getenv("BACKEND_PORT")

	connStr := "postgresql://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName + "?sslmode=" + dbSsl
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select version()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var version string
	for rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("version=%s\n", version)

	// Setup a Router
	r := chi.NewRouter()

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Println("Starting Server on :", backendPort)
	http.ListenAndServe(":"+backendPort, r)
}
