package http

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)
type ClientServer struct {
	Addr 			string
	Handler 		*mux.Router
	ReadTimeOut	   	time.Time
	WriteTimeOit	 time.Time
}

func NewServer(r *mux.Router, port string) *http.Server {
	return &http.Server{
		Addr:         port,          // configure the bind address
		Handler:      r,                // set the default handler
		ReadTimeout:  5 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
	}
}

// Run - start server.
func StartServer(s *http.Server) {
	log.Println("Starting server on port", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

	Shutdown(s)
}

func SetUpSubRouter(){

}
// Shutdown - shutdown server after get signal - "Interrupt"
func Shutdown(s *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)
	defer close(c)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx,_ := context.WithTimeout(context.Background(), 30*time.Second)

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("http server Shutdown: %v", err)
	}
}