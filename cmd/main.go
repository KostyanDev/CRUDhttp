package main

import (
	server "github.com/KostyanDev/CRUDhttp/helper"
	"github.com/gorilla/mux"

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

	// Start server
	go server.StartServer(s)

}
