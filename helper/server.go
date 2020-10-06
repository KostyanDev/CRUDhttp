package helper

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func NewServer(r *mux.Route, port string) *http.Server {
	return &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func StartServer(s *http.Server) {
	log.Println("Server started, port :", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error when server starting: %s\n", err)
		os.Exit(1)
	}
}
