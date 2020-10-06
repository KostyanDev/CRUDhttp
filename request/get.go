package request

import (
	"encoding/json"
	"fmt"

	"github.com/KostyanDev/CRUDhttp/model"
	"log"
	"net/http"
)

func GetListCars(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json")
	if len(model.GetToDoList()) == 0 {
		http.Error(w, "List of cars is empty.", http.StatusBadRequest)
	} else {
		list, err := json.MarshalIndent(model.GetToDoList(), "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(list))
	}
}
