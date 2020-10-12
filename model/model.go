package model

type STRToDoList struct {
	Id     int64    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

//type Task struct {
//	Text string `json:"name"`
//	Status *Status `json:status`
//}
//
//type Status struct {
//	status string `json:status`
//}
