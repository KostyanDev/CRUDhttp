package request

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"

	helper "github.com/KostyanDev/CRUDhttp/helper"
	"log"
	"net/http"
)

func HandlerGetToDoList(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json")
	if len(helper.GetToDoList()) == 0 {
		http.Error(w, "List of cars is empty.", http.StatusBadRequest)
	} else {
		list, err := json.MarshalIndent(helper.GetToDoList(), "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(list))
	}
}

func HandlerGetTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID parsing error.", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if list, ok := helper.GetTask(id); ok {
		c, err := json.MarshalIndent(list, "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(c))
	} else {
		http.Error(w, "Not Found.", http.StatusBadRequest)
	}
}

