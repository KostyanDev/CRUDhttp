package main

import (
	server "github.com/KostyanDev/CRUDhttp/internal/helper"
	handler "github.com/KostyanDev/CRUDhttp/request"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var port string

func init() {
	port = os.Getenv("PORT")

	if port == "" {
		panic("You need set port in .env file")
	}
}

func main() {

	r := mux.NewRouter()

	// GET
	r.HandleFunc("/tasks", handler.HandlerGetToDoList).Methods(http.MethodGet)
	r.HandleFunc("/task/{id:[0-9]+}", handler.HandlerGetTask).Methods(http.MethodGet)

	//POST
	r.HandleFunc("/task", handler.HandlerCreateTodoList).Methods(http.MethodPost, http.MethodGet)

	//UPDATE
	r.HandleFunc("/tasks/{id:[0-9]+}", handler.HandlerUpdateTask).Methods(http.MethodPut)
	//
	//DELETE
	r.HandleFunc("/tasks/{id:[0-9]+}", handler.HandlerDeleteList).Methods(http.MethodDelete)


	s := server.NewServer(r, ":"+port)
	// Start server
	server.StartServer(s)

}

