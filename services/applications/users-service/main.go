package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/TineoC/users-service/internal/handlers"
	"github.com/TineoC/users-service/internal/repositories"
	_ "github.com/lib/pq"
)

type Server struct {
	server *http.Server
	port   string
}

func NewServer() *Server {
	port := getPort()
	connStr := getConnectionString()

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("could not connect to DB: %v", err)
		exit()
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB: %v", err)
		exit()
	}

	userRepo := repositories.NewPostgresUserRepository(db)
	userService := handlers.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleHealth)
	mux.HandleFunc("/health", handlers.HandleHealth)

	mux.Handle("/users", userHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return &Server{
		server: server,
		port:   port,
	}
}

func (s *Server) Start() error {
	log.Printf("Starting Users Service on port %s", s.port)
	return s.server.ListenAndServe()
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("PORT not set, defaulting to %s", port)
	} else {
		log.Printf("Using PORT=%s from environment", port)
	}
	return port
}

func getConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Println("Database environment variables are required")
		exit()
	}

	// lib/pq expects keyword DSN, not URL
	return "host=" + host +
		" port=" + port +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" sslmode=disable"
}

func exit() {
	os.Exit(1)
}

func main() {
	server := NewServer()
	server.Start()
}
