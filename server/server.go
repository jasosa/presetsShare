package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jasosa/PresetsManagement/pkg/service"
)

// Server ...
type Server struct {
	service service.PresetsService
}

// NewServer initializes a new server
func NewServer() (*Server, error) {
	return nil, nil
}

// Start initializes Server
func (s *Server) Start() {
	s.initRoutes()
	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) initRoutes() {
	http.HandleFunc("/list", listPresetsHandler())
	http.HandleFunc("/show", showPresetHandler())
	http.HandleFunc("/upload", uploadPresethandler())
}

func listPresetsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("List")
	}
}

func showPresetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Show")
	}
}

func uploadPresethandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Upload")
	}
}
