package helper

import (
	"encoding/json"
	"io"
	"sync"
	model "github.com/KostyanDev/CRUDhttp/model"
)

// ToJSON serializes.
func (list *model.STRToDoList) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(list)
}

// FromJSON deserializes.
func (list *model.STRToDoList) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(list)
}

// GET methods
func GetToDoList() []*model.STRToDoList {
	rwm.Lock()
	defer rwm.Unlock()
	return toDoLists
}

func GetTask(id int64) (*model.STRToDoList, bool) {
	rwm.Lock()
	defer rwm.Unlock()
	for _,list := range toDoLists{
		if list.id == id {
			return list, true
		}
	}
	return nil, false
}



// UPDATE
func UpdateList(id int64, list *model.STRToDoList) bool {
	rwm.Lock()
	defer rwm.Unlock()
	for _,task := range toDoLists{
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

func AddToDoList(list *model.STRToDoList) bool {
	rwm.Lock()
	defer rwm.Unlock()
	lastID := int64(len(toDoLists) + 1)

	if list.Name == "" || list.Status == "" {
		return false
	}

	toDoLists[lastID] = list
	return true
}

// Delete methods

func DeleteCar(id int64) bool {
	rwm.Lock()
	defer rwm.Unlock()
	isFind := false
	for i, task := range toDoLists {
		if task.id == id {
			if i == len(toDoLists)-1 {
				toDoLists[i] = nil
				toDoLists = toDoLists[:i]
			} else {
				toDoLists = append(toDoLists[:i], toDoLists[i+1:]...)
			}
			isFind = true
			break
		}
	}

	if !isFind {
		return false
	}

	var index int64 = 1
	for _, task := range toDoLists {
		if task.id != index {
			task.id = index
			index++
		} else {
			index++
		}
	}

	return true
}
