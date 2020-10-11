package request

import (
	"fmt"
	"net/http"
	"github.com/KostyanDev/CRUDhttp/helper"
	"os"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		panic("Set port!")
	}
}


func HandlerCreateTodoList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "http://localhost:"+port+"/", http.StatusMovedPermanently)
	}

	var list helper.STRToDoList
	if r.Method == http.MethodPost {
		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			name := r.FormValue("name")
			status := r.FormValue("status")
			list = helper.STRToDoList{
				Name: name,
				Status: status,
			}
			if ok := helper.AddToDoList(&list); !ok {
				http.Error(w, "Error adding a sell order.", http.StatusBadRequest)
				return
			}
			fmt.Fprint(w, "(POST) SUCCESS! Added new task .")
		}else {
			if err := list.FromJSON(r.Body); err != nil {
				http.Error(w, "Error retrieving data from JSON.", http.StatusBadRequest)
				return
			}

			if ok := helper.AddToDoList(&list); !ok {
				http.Error(w, "Error adding a sell order.", http.StatusBadRequest)
				return
			}

			fmt.Fprint(w, "New task added.")
		}
	}
}