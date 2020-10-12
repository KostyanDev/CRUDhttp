package helper

import (
	model "github.com/KostyanDev/CRUDhttp/model"
)
type STRToDoList model.STRToDoList

var (
	ToDoLists = []*STRToDoList{
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
)
