package helper

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func NewServer(r *mux.Router, port string) *http.Server {
	return &http.Server{
		Addr:         port,             // configure the bind address
		Handler:      r,                // set the default handler
		ReadTimeout:  5 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
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
