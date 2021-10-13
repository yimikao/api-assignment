package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yimikao/api-assignment/api/config"
)

type Server struct {
	Router *mux.Router
}

// Initialize the router and database
func (s *Server) InitializeRouter() (*Server, string) {
	s.Router = mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8000"
	}
	config.InitDB(config.GetConfig())
	return s, port
}

// Run the app on it's router
func (s *Server) Run(host string) {

	log.Printf("Server up and running at http://localhost%s", host)
	log.Fatal(http.ListenAndServe(host, s.Router))
}
