package main

import (
	"gantt-server/models"
	"gantt-server/routers"
	"net/http"
)

func init() {
	if err := models.NewEngine("./data/task.db"); err != nil {
		panic(err)
	}
}

func main() {
	//server := http.Server{}
	http.Handle("/date/list", routers.HandleGetDateList())
	http.Handle("/date/add", routers.HandleAddDate())
	http.Handle("/date/update", routers.HandleUpdateDate())

	http.Handle("/task/list", routers.HandleGetTaskList())
	http.Handle("/task/add", routers.HandleAddTask())
	http.Handle("/task/update", routers.HandleUpdateTask())
	http.ListenAndServe("127.0.0.1:8080", nil)
}
