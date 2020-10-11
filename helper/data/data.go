package helper

import (
	"sync"
	model "github.com/KostyanDev/CRUDhttp/model"
)


var (
	toDoLists = []*model.STRToDoList{
		&model.STRToDoList{
			id: 1,
			Name:   "Move to shop",
			Status: "New",
		},
		&model.STRToDoList{
			id: 2,
			Name:   "Buy present for birthday",
			Status: "New",
		},
	}
	rwm sync.Mutex
)
