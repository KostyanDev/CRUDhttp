package request

import (
	"fmt"
	 helper "github.com/KostyanDev/CRUDhttp/helper"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func HandlerUpdateTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	var list helper.STRToDoList
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID parsing error.", http.StatusBadRequest)
		return
	}

	if err := list.FromJSON(r.Body); err != nil {
		http.Error(w, "Error retrieving data from JSON.", http.StatusBadRequest)
		return
	}

	if ok := helper.UpdateList(int64(id), &list); !ok {
		http.Error(w, "Data update error.", http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, " №%v task updated!", id)
	}


}


