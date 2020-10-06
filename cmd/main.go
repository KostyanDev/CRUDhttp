package main

import (
	"net/http";
	"os";
	"github.com/gorilla/mux"
	server "github.com/KostyanDev/CRUDhttp/helper"
)
var port string

func init(){
	port = os.Getenv("PORT")

	if port == ""{
		panic("You need set port in .env file")
	}
}

func main(){
	//r := mux.NewRouter()

	//s := server.NewServer()


}