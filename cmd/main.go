package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/kairavkkp/dedoxify-backend/routes"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	backendPort := os.Getenv("BACKEND_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSsl := os.Getenv("DB_SSL")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=" + dbSsl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established successfully!")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get underlying database connection:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database connection is not active:", err)
	}

	log.Println("Database connection is alive and active!")

	// Routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/public", routes.PublicRouter(db))
	r.Mount("/private", routes.PrivateRouter(db))

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	})
	log.Println("Starting Server on :", backendPort)
	http.ListenAndServe(":"+backendPort, r)
}
