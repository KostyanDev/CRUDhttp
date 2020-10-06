package data

import (
	"encoding/json"
	model "github.com/KostyanDev/CRUDhttp/model"
	"io"
	"sync"
)

type ToDoList struct {
	Name   string `json:"brand"`
	Status bool   `json: "status"`
}

var (
	toDoLists = map[uint64]*ToDoList{
		1: {
			Name:   "Move to shop",
			Status: false,
		},
		2: {
			Name:   "Buy present for birthday",
			Status: false,
		},
	}
	rwm sync.Mutex
)

// ToJSON serializes.
func (list *ToDoList) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(list)
}

// FromJSON deserializes.
func (list *ToDoList) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(list)
}

func GetToDoList() map[uint64]*ToDoList {
	rwm.Lock()
	defer rwm.Unlock()
	return toDoLists
}
