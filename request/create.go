package request

import (
	"fmt"
	"net/http"
	"github.com/KostyanDev/CRUDhttp/internal/helper"
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
	if err := list.FromJSON(r.Body); err != nil {
		http.Error(w, "Error JSON data ", http.StatusBadRequest)
		return
	}

	if ok := helper.AddToDoList(&list); !ok {
		http.Error(w, "Error adding a sell order.", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "New task added.")
}