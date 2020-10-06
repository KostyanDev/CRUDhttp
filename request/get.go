package request

import (
	"encoding/json"
	"fmt"

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
