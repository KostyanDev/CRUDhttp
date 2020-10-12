package helper

import (
	"encoding/json"
	"io"
	model "github.com/KostyanDev/CRUDhttp/model"
	data "github.com/KostyanDev/CRUDhttp/helper/data"
	"sync"
)

type STRToDoList model.STRToDoList

var (
	rwm sync.Mutex
	lastID int = 3
	todoLists = data.ToDoList
)

// ToJSON serializes.
func (list *STRToDoList) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(list)
}

// FromJSON deserializes.
func (list *STRToDoList) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(list)
}

// GET methods
func GetToDoList() []*STRToDoList {
	rwm.Lock()
	defer rwm.Unlock()
	return todoLists
}

func GetTask(id int64) (*STRToDoList, bool) {
	rwm.Lock()
	defer rwm.Unlock()
	for _,list := range todoLists{
		if list.id == id {
			return list, true
		}
	}
	return nil, false
}



// UPDATE
func UpdateList(id int64, list *STRToDoList) bool {
	rwm.Lock()
	defer rwm.Unlock()
	for _,task := range todoLists{
		if task.id == id {
			if len(list.Name) != 0 || list.Name != ""{
				task.Name = list.Name
			}
			if len(list.Status) != 0 || list.Status != ""{
				task.Status = list.Status
			}
			return true
		}
	}
	return false
}

// POST methods

func AddToDoList(list *STRToDoList) bool {
	rwm.Lock()
	defer rwm.Unlock()
	lastID := int64(len(todoLists) + 1)

	if list.Name == "" || list.Status == "" {
		return false
	}

	todoLists[lastID] = list
	return true
}

// Delete methods

func DeleteCar(id int64) bool {
	rwm.Lock()
	defer rwm.Unlock()
	isFind := false
	for i, task := range todoLists {
		if task.id == id {
			if i == len(todoLists)-1 {
				todoLists[i] = nil
				todoLists = todoLists[:i]
			} else {
				todoLists = append(todoLists[:i], todoLists[i+1:]...)
			}
			isFind = true
			break
		}
	}

	if !isFind {
		return false
	}

	var index int64 = 1
	for _, task := range todoLists {
		if task.id != index {
			task.id = index
			index++
		} else {
			index++
		}
	}

	return true
}
