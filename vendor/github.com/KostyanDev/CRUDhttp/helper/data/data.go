package helper

import (
	model "github.com/KostyanDev/CRUDhttp/model"
)
type STRToDoList model.STRToDoList

var (
	ToDoLists = []*STRToDoList{
		&STRToDoList{
			id: 1,
			Name:   "Move to shop",
			Status: "New",
		},
		&STRToDoList{
			id: 2,
			Name:   "Buy present for birthday",
			Status: "New",
		},
	}
)
