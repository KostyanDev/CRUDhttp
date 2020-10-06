package request

import (
	"fmt"
	"net/http"
	helper "github.com/KostyanDev/CRUDhttp/helper"
)

func CreateTodoList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "http://localhost:"+port+"/", http.StatusMovedPermanently)
	}

	var list helper.TodoList
	if r.Method == http.MethodPost {
		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			name := r.FormValue("name")
			status := r.FormValue("status")
			list = helper.ToDoList{
				name: name,
				status: status,
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

			if ok := list.AddToDoList(&list); !ok {
				http.Error(w, "Error adding a sell order.", http.StatusBadRequest)
				return
			}

			fmt.Fprint(w, "(JSON) SUCCESS! Added new car sale announcement.")
		}
	}
}