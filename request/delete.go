package request

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/KostyanDev/CRUDhttp/helper"
)

func HandlerDeleteList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Error in ParseUint.", http.StatusInternalServerError)
	}

	if ok := helper.DeleteToDoList(id); !ok {
		http.Error(w, "Not found or was deleted.", http.StatusNotFound)
	} else {
		fmt.Fprintf(w, " Task №%v deleted", id)
	}
}