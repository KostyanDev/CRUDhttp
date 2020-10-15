package helper

import (
	"encoding/json"
	"io"
	model "github.com/KostyanDev/CRUDhttp/api/model"

	"sync"
)

type STRToDoList model.STRToDoList

var todoLists = []*STRToDoList{
	&STRToDoList{
		Id: 1,
		Name:   "Move to shop",
		Status: "New",
	},
	&STRToDoList{
		Id: 2,
		Name:   "Buy present for birthday",
		Status: "New",
	},
}

var (
	rwm sync.Mutex
	lastID int64 = 3
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
		if list.Id == id {
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
		if task.Id == id {
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
	if len(todoLists) > 0 {
		lastID = todoLists[len(todoLists)-1].Id + 1
	} else {
		lastID = 1
	}

	if list.Name == "" || list.Status == "" {
		return false
	}

	list.Id = lastID
	todoLists = append(todoLists, list)
	return true
}

// Delete methods

func DeleteToDoList(id int64) bool {
	rwm.Lock()
	defer rwm.Unlock()
	isFind := false
	for i, task := range todoLists {
		if task.Id == id {
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
		if task.Id != index {
			task.Id = index
			index++
		} else {
			index++
		}
	}

	return true
}
