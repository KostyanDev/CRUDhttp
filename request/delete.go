package request

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	helper "github.com/KostyanDev/CRUDhttp/helper"
)

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Error in ParseUint.", http.StatusInternalServerError)
	}

	if ok := helper.DeleteCar(id); !ok {
		http.Error(w, "Not found or was deleted.", http.StatusNotFound)
	} else {
		fmt.Fprintf(w, "(JSON) SUCCESS! Deleted ID = %v", id)
	}
}