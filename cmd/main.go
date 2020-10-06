package main

import (
	server "github.com/KostyanDev/CRUDhttp/helper"
	handler "github.com/KostyanDev/CRUDhttp/request"
	"github.com/gorilla/mux"
	"net/http"

	//"net/http";
	"os"
	//_ "github.com/gorilla/mux"
)

var port string

func init() {
	port = os.Getenv("PORT")

	if port == "" {
		panic("You need set port in .env file")
	}
}

func main() {

	rout := mux.NewRouter()

	s := server.NewServer(rout, ":"+port)
	// GET
	rout.HandleFunc("/tasks", handler.HandlerGetToDoList).Methods(http.MethodGet)

	// Start server
	go server.StartServer(s)

}
