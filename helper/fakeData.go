package helper

import (
	"encoding/json"
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

// GET methods
func GetToDoList() map[uint64]*ToDoList {
	rwm.Lock()
	defer rwm.Unlock()
	return toDoLists
}

// UPDATE
func UpdateCar(id uint64, list *ToDoList) bool {
	rwm.Lock()
	defer rwm.Unlock()
	_, ok := toDoLists[id]
	if !ok {
		return false
	}

	toDoLists[id] = list
	return true
}

// POST methods

func AddToDoList(list *ToDoList) bool {
	rwm.Lock()
	defer rwm.Unlock()
	lastID := uint64(len(toDoLists) + 1)

	if list.Name == "" || list.Status == nil {
		return false
	}

	toDoLists[lastID] = list
	return true
}

// Delete methods

func DeleteToDoList(id uint64) bool {
	rwm.Lock()
	defer rwm.Unlock()
	if _, ok := toDoLists[id]; !ok {
		return false
	}
	delete(toDoLists, id)
	return true
}
