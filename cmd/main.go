package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kairavkkp/dedoxify-backend/routes"
	_ "github.com/lib/pq"
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

	// Routes
	router := routes.SetupRouter()

	log.Println("Starting Server on :", backendPort)
	http.ListenAndServe(":"+backendPort, router)
}
